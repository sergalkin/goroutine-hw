package concurency

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

var taskError = errors.New("worker aborted due to task error")

func Concurrency(tasks []func() error, maxWorkers int, maxErrors int) error {

	workersCh := make(chan func() error, maxWorkers)
	errorCh := make(chan error, maxWorkers)
	defer close(errorCh)

	shutdownCh := make(chan bool)
	errCnt := &errorsCnt{errorThreshold: maxErrors}

	go func() {
		for err := range errorCh {
			errCnt.Inc()
			fmt.Println("Detected an error:", err)
			if errCnt.isThresholdReached() {
				fmt.Println("Threshold number of errors has reached. Aborting...")
				close(shutdownCh)
				return
			}
		}
	}()

	eg := errgroup.Group{}

	for i := 1; i <= maxWorkers; i++ {
		i := i
		eg.Go(func() error {
			for task := range workersCh {
				select {
				case <-shutdownCh:
					fmt.Println("Worker aborted...", i)
					return taskError
				default:
					fmt.Println("Worker started...", i)
					if err := task(); err != nil {
						errorCh <- err
					}
					fmt.Println("Worker stopped", i)
				}
			}
			fmt.Println("Worker finished...", i)
			return nil
		})
	}

	for _, task := range tasks {
		select {
		case <-shutdownCh:
			break
		default:
			workersCh <- task
		}
	}
	close(workersCh)

	err := eg.Wait()
	if err != nil {
		return err
	} else if errCnt.Get() != 0 {
		return taskError
	}
	return nil
}
