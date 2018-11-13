package scheduler

import "crawler/concurrent/engine"

type SimpleScheduler struct{
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler)Submit(r engine.Request) {
	  s.WorkerChan <- r
}

func (s *SimpleScheduler)ConfigureMasterWorkerChan(r chan engine.Request) {
	  s.WorkerChan = r
}