package taskrunner

import (
	"log"
	"time"
)

// Worker for work
type Worker struct {
	// ? ptr ?
	ticker *time.Ticker
	runner *Runner
}

// NewWorker add new worker
func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWork() {
	for {
		select {
		case <-w.ticker.C:
			log.Println("Starting to work ...")
			go w.runner.StartAll()
		}
	}
}

// Start start all work
func Start() {
	// ? where is the data chan of the last two method ?
	// yeah the contrustor is type is a func
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	// 10 seconds
	w := NewWorker(10, r)
	go w.startWork()
}
