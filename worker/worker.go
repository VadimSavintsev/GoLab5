package worker

import (
	"fmt"

	"github.com/VadimSavintsev/GoLab5/mutex"
)

func RunWorkers(count int) {
	m := mutex.New(count)

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer m.Unlock()
			fmt.Printf("Goroutine %d: Hello, World\n", id)
		}(i)
	}

	m.Wait() // Wait for all goroutines to finish
}
