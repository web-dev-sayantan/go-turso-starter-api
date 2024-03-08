package playground

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
	Concurrency int
	TasksChan   chan Task
	Wg          *sync.WaitGroup
}

func (w *WorkerPool) worker() {
	for task := range w.TasksChan {
		fmt.Printf("Task type: %T\n", task)
		if task == nil {
			fmt.Println("Task is nil")
		}
		task.Process()
		w.Wg.Done()
	}
}

// Start function initializes the worker pool and executes the tasks in parallel.
//
// No parameters.
// No return values.
func (w *WorkerPool) Start() {
	fmt.Println("Starting worker pool with ", w.Concurrency, " workers")
	w.Wg = &sync.WaitGroup{}
	w.TasksChan = make(chan Task, len(*w.Tasks))
	for i := 0; i < w.Concurrency; i++ {
		go w.worker()
	}
	fmt.Printf("Adding counter to waitgroup %d\n", len(*w.Tasks))
	w.Wg.Add(len(*w.Tasks))
	for _, task := range *w.Tasks {
		w.TasksChan <- task
	}
	fmt.Printf("Closing tasks channel with %d tasks\n", len(*w.Tasks))
	defer close(w.TasksChan)
	w.Wg.Wait()
}
