package semaphore

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semCount chan  struct{}
}

func  NewSemaphore(permits int) Semaphore {
	return &semaphore{semCount: make(chan struct{}, permits)}
}

func (s *semaphore) Acquire() {
	s.semCount<- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semCount
}
