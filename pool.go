package goroutine_pool

import "log"

type Pool struct {
	EntranceChan chan *Task // 对外开放的Task入口

	jobsChan chan *Task // 内部的Task队列

	workSize int // 协程池最大的task数量
}

func NewPool(size int) *Pool {
	return &Pool{
		EntranceChan: make(chan *Task),
		jobsChan:     make(chan *Task),
		workSize:     size,
	}
}

// 创建一个worker，并且让这个worker去工作
func (p *Pool) worker(workerId int) {
	for {
		select {
		case task := <-p.jobsChan:
			task.Execute()
			log.Println("worker ID:", workerId, " 执行了一个任务")
		}
	}
}

// 启动线程池
func (p *Pool) run() {
	for i := 0; i < p.workSize; i++ {
		go p.worker(i)
	}
	for {
		select {
		case task := <-p.EntranceChan:
			p.jobsChan <- task
		}
	}
}
