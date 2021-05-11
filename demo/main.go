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

		go func(v int) {
			defer semaphore.Release()
			longRunningProcess(v)
			if v == proccess {
				doneS <- true
			}
		}(i)
	}

	<-doneS
}

func longRunningProcess(ID int) {
	fmt.Println(time.Now().Format("15:04:05.000"), "Running task with ID", ID)
	time.Sleep(5 * time.Second)
}
