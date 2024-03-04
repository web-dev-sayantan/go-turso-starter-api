package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Process()
}

type EmailTask struct {
	Email   string
	Subject string
	Message string
}

func (e EmailTask) Process() {
	fmt.Println("Sending email to ", e.Email)
	time.Sleep(2 * time.Second)
}

type ImageProcessingTask struct {
	ImagePath string
	Width     int
	Height    int
}

func (i ImageProcessingTask) Process() {
	fmt.Println("Processing image ", i.ImagePath)
	time.Sleep(4 * time.Second)
	fmt.Println("Resizing image to ", i.Width, "x", i.Height)
	time.Sleep(1 * time.Second)
}

type WorkerPool struct {
	Tasks       *[]Task
	concurrency int
	tasksChan   chan Task
	wg          *sync.WaitGroup
}

func (w *WorkerPool) worker() {
	for task := range w.tasksChan {
		fmt.Printf("Task type: %T\n", task)
		if task == nil {
			fmt.Println("Task is nil")
		}
		task.Process()
		w.wg.Done()
	}
}

// Start function initializes the worker pool and executes the tasks in parallel.
//
// No parameters.
// No return values.
func (w *WorkerPool) Start() {
	fmt.Println("Starting worker pool with ", w.concurrency, " workers")
	w.wg = &sync.WaitGroup{}
	w.tasksChan = make(chan Task, len(*w.Tasks))
	for i := 0; i < w.concurrency; i++ {
		go w.worker()
	}
	fmt.Printf("Adding counter to waitgroup %d\n", len(*w.Tasks))
	w.wg.Add(len(*w.Tasks))
	for _, task := range *w.Tasks {
		w.tasksChan <- task
	}
	fmt.Printf("Closing tasks channel with %d tasks\n", len(*w.Tasks))
	defer close(w.tasksChan)
	w.wg.Wait()
}

func main() {
	start := time.Now()
	tasks := make([]Task, 20)
	for i := 0; i < 10; i++ {
		tasks[i] = &ImageProcessingTask{
			ImagePath: "image.png",
			Width:     100,
			Height:    100,
		}
	}
	for i := 10; i < 20; i++ {
		tasks[i] = &EmailTask{
			Email:   "a@a.com",
			Subject: "Subject",
			Message: "Message",
		}
	}
	workerPool := WorkerPool{
		Tasks:       &tasks,
		concurrency: 6,
	}
	workerPool.Start()
	fmt.Println("Time taken to process: ", time.Since(start))
}
