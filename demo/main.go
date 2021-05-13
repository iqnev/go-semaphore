package main

import (
	"fmt"
	"time"

	"github.com/iqnev/go-semaphore"
)

func main() {

	semaphore := semaphore.NewSemaphore(2)

	doneS := make(chan bool, 1)

	proccess := 13

	for i := 1; i <= proccess; i++ {
		semaphore.Acquire()

		go func(p int) {
			defer semaphore.Release()
			runningProcess(p)
			if p == proccess {
				doneS <- true
			}
		}(i)
	}

	<-doneS
}

func runningProcess(ID int) {
	fmt.Println(time.Now().Format("15:04:05.000"), "Running task with ID", ID)
	time.Sleep(5 * time.Second)
}
