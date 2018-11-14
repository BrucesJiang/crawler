package scheduler

import (
	"crawler/concurrent/engine"
)

type QueuedScheduler struct {
	RequestChan chan engine.Request
	WorkerChan chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.RequestChan <- r
}

func (s *QueuedScheduler) WorkerReady(r chan engine.Request) {
	s.WorkerChan <- r
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {

}


func (s *QueuedScheduler) Run() {
	//初始化参量
	s.RequestChan = make(chan engine.Request)
	s.WorkerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select{
			case r := <- s.RequestChan:
				requestQ = append(requestQ, r)
			case w := <- s.WorkerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}