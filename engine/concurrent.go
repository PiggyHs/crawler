package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfiglerMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.ConfiglerMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Go Item #%d %s", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := Worker(request)

			if err != nil {
				continue
			}
			out <- result
		}
	}()
}