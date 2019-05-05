package scheduler

import "crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (e *QueueScheduler) Submit(r engine.Request) {
	e.requestChan <- r
}

func (e *QueueScheduler) WorkerReady(
	w chan engine.Request) {
	e.workerChan <- w
}

func (e *QueueScheduler) ConfiglerMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (e *QueueScheduler) Run() {
	e.requestChan = make(chan engine.Request)
	e.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 &&
				len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case r := <-e.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-e.workerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
