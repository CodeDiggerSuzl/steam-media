package taskrunner

import (
	"log"
	"time"
)

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
			log.Println("start Work ...")
			go w.runner.StartAll()
		}
	}
}

func Start() {
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWork()
}
