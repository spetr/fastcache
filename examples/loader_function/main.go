package main

import (
	"fmt"

	"github.com/spetr/fastcache"
)

func main() {
	gc := fastcache.New(10).
		LFU().
		LoaderFunc(func(key interface{}) (interface{}, error) {
			return fmt.Sprintf("%v-value", key), nil
		}).
		Build()

	v, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
