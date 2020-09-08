package main

import (
	"errors"
	"fmt"
	"github.com/sergalkin/otus-hw-5/internal/concurency"
	"math/rand"
	"time"
)

func main() {
	fns := []func() error {task, task, task, task}
	fmt.Println("Dispatcher Started...")
	err := concurency.Concurrency(fns, 2, 1)
	if err != nil {
		fmt.Println("Dispatcher finished with errors", err)
	}
	fmt.Println("Dispatcher ended...", err)
}

func task() error {
	fmt.Println("Starting task...")
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if rand.Intn(10) >= 5 {
		fmt.Println("Time since start to get an error", time.Since(start))
		return errors.New("task has an error during execution")
	}
	fmt.Println("Ending successfully... Time since start", time.Since(start))
	return nil
}