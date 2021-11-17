package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
)

func main() {
	// Background type of empty context
	bg := context.Background()

	// TODO type of empty context
	empty := context.TODO()

	fmt.Printf("hash 1: %s\n", process(bg, "golangislife"))

	fmt.Printf("hash 2: %s", process(empty, "golangislife"))
}

func process(ctx context.Context, username string) string {
	// Value call on empty context is always nil
	val := ctx.Value("passwd")
	if val != nil {
		log.Fatalf("expected value [%v] to be nil", val)
	}

	// other interface method calls return nils as well

	return fmt.Sprintf("%v", md5.Sum([]byte(username)))
}
