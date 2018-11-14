package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WokerCount int
}

type Scheduler interface {
	 ReadyNotifier //组合的方式
	 Submit(Request)
	 WorkerChan() chan Request
	 Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seed ...Request) {
	out := make(chan ParseResult)
	//in := make(chan  Request)
	//c.Scheduler.ConfigureMasterWorkerChan(in)
	c.Scheduler.Run() //在这里创建 in

	//构建多少个Worker, 这个事情应该由Scheduler决定
	for i := 0;  i < c.WokerCount; i ++ {
		createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	//等待初始化完毕
	for _, r := range seed {
		//向调度器提交种子请求
		c.Scheduler.Submit(r)
	}

	//接收Worker处理的结果
	for {
		result := <- out

		for _, item := range result.Items {
			log.Printf("Got item:  %v\n", item)
		}

		//将新的请求放到调度器
		for _, r := range result.Requests {
			c.Scheduler.Submit(r)
		}
	}
}

//通过 goroutine创建一个新的Worker工作, 公用一个chan
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	//每个Worker都有一个chan
	//in := make(chan Request)
	go func() {
		for {
			//tell scheduler I'm ready
			ready.WorkerReady(in)
			request := <- in
			result, err := Worker(request)
			if  err != nil {
				continue
			}
			out <- result
		}
	}()
}