package threadpool

import "sync"

type Pool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

func New(size int) *Pool {
	p := &Pool{tasks: make(chan func(), size)}
	for i := 0; i < size; i++ {
		go p.worker()
	}
	return p
}

func (p *Pool) worker() {
	for task := range p.tasks {
		task()
		p.wg.Done()
	}
}

func (p *Pool) Add(task func()) {
	p.wg.Add(1)
	p.tasks <- task
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) Close() {
	close(p.tasks)
}
