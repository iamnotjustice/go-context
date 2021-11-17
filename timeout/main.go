package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Background type of empty context as a base for our contexts
	base := context.Background()

	// Timeout context as a child of base, and cancel as CancelFunc which we can use to send
	// cancellation message
	withTimeout, cancel := context.WithTimeout(base, time.Second*2)

	go process(withTimeout, 10)

	// we do not need to stop our operation manually now, it'll cancel itself after timeout
	// but it's a good thing to call cancel eventually, just so the context is cancelled for sure
	defer cancel()

	time.Sleep(time.Second * 4)
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
