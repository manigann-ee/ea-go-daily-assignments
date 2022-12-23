package executor

import (
	"errors"
	"sync"
)

type Service struct {
	jobChan  chan func()
	isClosed bool
}

func (s *Service) Start(workersCnt int, setup func()) {
	jobChan := make(chan func())
	s.jobChan = jobChan
	wg := sync.WaitGroup{}
	wg.Add(workersCnt)

	for i := 0; i < workersCnt; i++ {
		go worker(setup, &wg, jobChan)
	}

	wg.Wait()
}

func (s *Service) Run(job func()) error {
	if s.isClosed {
		return errors.New("closed")
	}
	s.jobChan <- job
	return nil
}

func (s *Service) Close() {
	close(s.jobChan)
	s.isClosed = true
}

func worker(setup func(), wg *sync.WaitGroup, jobChan chan func()) {
	setup()
	wg.Done()
	for {
		job, ok := <-jobChan
		if !ok {
			break
		}
		job()
	}
}
