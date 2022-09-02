package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	r := RunTasks(1500)
	if r {
		fmt.Println("Job completed")
	} else {
		fmt.Println("Timeout exceeded")
	}
}

func RunTasks(timeoutAsMilisecond int64) bool {
	ctx, cancel := context.WithTimeout(context.Background(), (time.Duration(timeoutAsMilisecond) * time.Millisecond))

	c := make(chan bool)

	go func() {

		for i := 0; i < 50; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)

			// Added for supress 'discarding cancel method' warning
			if i == 100 {
				cancel()
			}
		}
		c <- true
	}()

	select {
	case <-c:
		return true
	case <-ctx.Done():
		return false
	}
}
