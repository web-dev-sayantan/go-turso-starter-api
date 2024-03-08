package main

import (
	"fmt"
	"time"

	"github.com/ishanz23/go-turso-starter-api/playground"
)

func main() {
	start := time.Now()
	tasks := make([]playground.Task, 20)
	for i := 0; i < 10; i++ {
		tasks[i] = &playground.ImageProcessingTask{
			ImagePath: "image.png",
			Width:     100,
			Height:    100,
		}
	}
	for i := 10; i < 20; i++ {
		tasks[i] = &playground.EmailTask{
			Email:   "a@a.com",
			Subject: "Subject",
			Message: "Message",
		}
	}
	workerPool := playground.WorkerPool{
		Tasks:       &tasks,
		Concurrency: 6,
	}
	workerPool.Start()
	fmt.Println("Time taken to process: ", time.Since(start))
}
