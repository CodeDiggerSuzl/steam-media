package taskrunner

import "log"

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	longLived  bool // to recycle or not
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longLived bool, d fn, e fn) *Runner {
	return &Runner{
		// with buffer channel
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longLived:  longLived,
		dataSize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) StartAll() {
	// if don't send something to the channel, it will stuck HERE
	r.Controller <- READY_TO_DISPATCH
	r.startDisPatch()
}

func (r *Runner) startDisPatch() {

	defer func() {
		// need to recycle
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}() // instant call the method

	for {
		// non-blocking
		select {
		// task
		case c := <-r.Controller:
			// ready to dispatch
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					log.Printf("Error during select READ_TO_DISPATCH: %v", err)
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}
			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		// if error happend
		case e := <-r.Error:
			if e == CLOSE {
				return // return the function
			}
		default:
		}
	}
}
