package semaphore

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semaphoreCount chan  struct{}
}

func  NewSemaphore(permits int) *semaphore {
	return &semaphore{semaphoreCount: make(chan struct{}, permits)}
}

func (s *semaphore) Acquire() {
	s.semaphoreCount<- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semaphoreCount
}
