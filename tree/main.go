package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"

	"github.com/pborman/uuid"
)

const (
	passwdKey      = ctxKey("password")
	currentTimeKey = ctxKey("current_time")
	processTypeKey = ctxKey("process_type")
	requestID      = ctxKey("req_id")
)

type ctxKey string

func main() {
	// Background type of empty context as a base for our contexts
	base := context.Background()

	// Context with Value as a child of base (background) context
	valueCtx := context.WithValue(base, passwdKey, "qwerty12345")

	go fmt.Printf("hash with pass: %s\n", process(context.WithValue(valueCtx, requestID, uuid.NewRandom()), "golangislife"))

	go fmt.Printf("hash no pass: %s\n", process(context.WithValue(base, requestID, uuid.NewRandom()), "golangislife"))

	time.Sleep(time.Second * 1)
}

func process(ctx context.Context, username string) string {
	// store the request start time in new context
	ctx = context.WithValue(ctx, currentTimeKey, time.Now().Format(time.RFC3339Nano))

	pass := ctx.Value(passwdKey)
	if pass != nil {
		ctx = context.WithValue(ctx, processTypeKey, "with_pass")
		return hash(ctx, fmt.Sprintf("%s:%s", pass, username))
	}

	ctx = context.WithValue(ctx, processTypeKey, "no_pass")
	return hash(ctx, username)
}

func hash(ctx context.Context, toHash string) string {
	fmt.Printf("DEBUG: request_id: %v, time_started: %v, type: %v, toHash: %s\n",
		ctx.Value(requestID),
		ctx.Value(currentTimeKey),
		ctx.Value(processTypeKey),
		toHash,
	)

	return fmt.Sprintf("%v", md5.Sum([]byte(toHash)))
}
