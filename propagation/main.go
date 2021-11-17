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
	// we start a side operation with the new cancel context
	newCtx, newCancel := context.WithCancel(ctx)
	go sideProcess(newCtx, nums*10)

	// always ensure that the context is cancelled at some point
	defer newCancel()

	i := 0
	for i < nums {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exited from main operation")
			fmt.Printf("cancelaltion reason for main operation is [%v]\n", ctx.Err())
			return
		default:
			fmt.Printf("Main Operation progress: [%d%%] done\n", i*10)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}

func sideProcess(ctx context.Context, nums int) {
	i := 0
	for i < nums {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exited from side Process")
			fmt.Printf("cancelaltion reason for side process is [%v]\n", ctx.Err())
			return
		default:
			fmt.Printf("Side operation progress: [%d%%] done\n", i)
			time.Sleep(time.Millisecond * 50)
			i++
		}
	}
}
