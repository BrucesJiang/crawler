package scheduler

import "crawler/concurrent/engine"

type SimpleScheduler struct{
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler)Submit(r engine.Request) {
	//如果这个会卡死
	//s.WorkerChan <- r
	//解决方案， 使用goroutine去通知,为每一个Request创建goroutine
	go func() {s.workerChan <- r}()
}
