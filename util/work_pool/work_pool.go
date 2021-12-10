package work_pool

import (
	"sync"
)

type WorkBase interface {
	Task(params interface{})
}

type WorkPool struct {
	ch chan WorkBase
	wg sync.WaitGroup
}

func NewWorkPool(size int) *WorkPool {
	pool := new(WorkPool)
	pool.ch = make(chan WorkBase, size)

	pool.wg.Add(size)
	for i := 0; i < size; i++ {
		go func(index int) {
			for msg := range pool.ch {
				msg.Task(index)
			}
			pool.wg.Done()
		}(i)
	}

	return pool
}

func (p *WorkPool) Run(task *WorkBase) {
	p.ch <- *task
}

func (p *WorkPool) Close() {
	close(p.ch)
}
