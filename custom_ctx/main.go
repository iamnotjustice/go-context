package main

import (
	"context"
	"fmt"
	"time"
)

type CustomCtx struct {
	a string
}

func (c *CustomCtx) Value(key interface{}) interface{} {
	return c.a
}

func (c *CustomCtx) Done() <-chan struct{} {
	return nil
}
func (c *CustomCtx) Deadline() (dl time.Time, ok bool) {
	return
}

func (c *CustomCtx) Err() error {
	return nil
}

func main() {
	val := av(&CustomCtx{a: "a"})

	fmt.Print(val)
}

func av(ctx context.Context) string {
	v, _ := ctx.Value("a").(string)

	return v
}
