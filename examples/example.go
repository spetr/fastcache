package main

import (
	"context"
	"fmt"

	"github.com/spetr/fastcache"
)

func main() {
	gc := fastcache.New(10).
		LFU().
		Build()
	gc.Set("key", "ok")

	v, err := gc.Get(context.Background(), "key")
	if err != nil {
		panic(err)
	}
	fmt.Println("value:", v)
}
