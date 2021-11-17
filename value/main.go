package main

import (
	"context"
	"crypto/md5"
	"fmt"
)

const (
	passwdKey = ctxKey("password")
)

type ctxKey string

func main() {
	// Background type of empty context as a base for our contexts
	base := context.Background()

	// Context with Value as a child of base (background) context
	valueCtx := context.WithValue(base, passwdKey, "qwerty12345")

	fmt.Printf("hash 1: %s\n", process(base, "golangislife"))

	// this shows that the hash is different, as a result of the passed value.
	fmt.Printf("hash 2: %s", process(valueCtx, "golangislife"))
}

func process(ctx context.Context, username string) string {
	// Value call on WithValue context returns interface{} value if key is present,
	// nil if it's not.
	pass := ctx.Value(passwdKey)
	if pass != nil {
		return fmt.Sprintf("%v", md5.Sum([]byte(fmt.Sprintf("%s:%s", username, pass))))
	}

	// other interface method calls still return nils

	return fmt.Sprintf("%v", md5.Sum([]byte(username)))
}
