package main

import (
	"context"
	"fmt"

	"github.com/spetr/fastcache"
)

func main() {
	gc := fastcache.New(10).
		LFU().
		LoaderFunc(func(_ context.Context, key interface{}) (interface{}, error) {
			return fmt.Sprintf("%v-value", key), nil
		}).
		Build()

	v, err := gc.Get(context.Background(), "key")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
