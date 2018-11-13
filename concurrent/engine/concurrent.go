package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WokerCount int
}

type Scheduler interface {
	 Submit(Request)
	 ConfigureMasterWorkerChan(chan Request)
}


func (c *ConcurrentEngine) Run(seed ...Request) {
	in := make(chan  Request)
	out := make(chan ParseResult)
	c.Scheduler.ConfigureMasterWorkerChan(in)

	//构建多少个Worker, 这个事情应该由Scheduler决定
	for i := 0;  i < c.WokerCount; i ++ {
		createWorker(in, out)
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

//通过 goroutine创建一个新的Worker工作
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <- in
			result, err := Worker(request)
			if  err != nil {
				continue
			}
			out <- result
		}
	}()
}