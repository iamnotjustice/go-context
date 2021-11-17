package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Background type of empty context as a base for our contexts
	base := context.Background()

	// Cancellation context as a child of base, and cancel as CancelFunc which we use to send
	// cancellation message
	withCancel, cancel := context.WithCancel(base)

	go process(withCancel, 10)

	// we stop our operation 2 seconds after we started
	time.Sleep(time.Second * 2)
	cancel()

	time.Sleep(time.Second * 2)
}

func process(ctx context.Context, nums int) {
	i := 0
	for i < nums {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Printf("cancelaltion reason is [%v]", ctx.Err())
			return
		default:
			fmt.Printf("Operation progress: [%d%%] done\n", i*10)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
