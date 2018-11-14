package scheduler

import "crawler/concurrent/engine"

type SimpleScheduler struct{
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler)Submit(r engine.Request) {
	//如果这个会卡死
	//s.WorkerChan <- r
	//解决方案， 使用goroutine去通知,为每一个Request创建goroutine
	go func() {s.WorkerChan <- r}()
}

func (s *SimpleScheduler)ConfigureMasterWorkerChan(r chan engine.Request) {
	  s.WorkerChan = r
}