package fastcache

import (
	"fmt"
	"testing"
	"time"
)

func TestLRUGet(t *testing.T) {
	size := 1000
	gc := buildTestCache(t, cacheTypeLRU, size)
	testSetCache(t, gc, size)
	testGetCache(t, gc, size)
}

func TestLoadingLRUGet(t *testing.T) {
	size := 1000
	gc := buildTestLoadingCache(t, cacheTypeLRU, size, loader)
	testGetCache(t, gc, size)
}

func TestLRULength(t *testing.T) {
	gc := buildTestLoadingCache(t, cacheTypeLRU, 1000, loader)
	gc.Get(ctx, "test1")
	gc.Get(ctx, "test2")
	length := gc.Len(true)
	expectedLength := 2
	if length != expectedLength {
		t.Errorf("Expected length is %v, not %v", length, expectedLength)
	}
}

func TestLRUEvictItem(t *testing.T) {
	cacheSize := 10
	numbers := 11
	gc := buildTestLoadingCache(t, cacheTypeLRU, cacheSize, loader)

	for i := 0; i < numbers; i++ {
		_, err := gc.Get(ctx, fmt.Sprintf("Key-%d", i))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestLRUGetIFPresent(t *testing.T) {
	testGetIFPresent(t, cacheTypeLRU)
}

func TestLRUHas(t *testing.T) {
	gc := buildTestLoadingCacheWithExpiration(t, cacheTypeLRU, 2, 10*time.Millisecond)

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			gc.Get(ctx, "test1")
			gc.Get(ctx, "test2")

			if gc.Has("test0") {
				t.Fatal("should not have test0")
			}
			if !gc.Has("test1") {
				t.Fatal("should have test1")
			}
			if !gc.Has("test2") {
				t.Fatal("should have test2")
			}

			time.Sleep(20 * time.Millisecond)

			if gc.Has("test0") {
				t.Fatal("should not have test0")
			}
			if gc.Has("test1") {
				t.Fatal("should not have test1")
			}
			if gc.Has("test2") {
				t.Fatal("should not have test2")
			}
		})
	}
}

func TestLRUCache_SetSize(t *testing.T) {
	gc := buildTestCache(t, cacheTypeLRU, 5)
	for i := 0; i < 10; i++ {
		err := gc.Set(fmt.Sprintf("Test_%d", i), i)
		if err != nil {
			t.Fatal("error setting in cache")
		}
	}
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			_ = gc.Set("test_11", 11)
			if gc.Len(false) > 5 {
				t.Fatalf("size %d is greater than expected size 5", gc.Len(false))
			}
			time.Sleep(time.Millisecond * 10)
			t.Logf("current size = %d", gc.Len(false))

			gc.SetSize(4)
			_ = gc.Set("test_12", 12)
			if gc.Len(false) > 4 {
				t.Fatalf("size %d is greater than expected size 4", gc.Len(false))
			}
			t.Logf("current size = %d", gc.Len(false))

			gc.SetSize(3)
			_ = gc.Set("test_13", 13)
			if gc.Len(false) > 3 {
				t.Fatalf("size %d is greater than expected size 3", gc.Len(false))
			}
			t.Logf("current size = %d", gc.Len(false))
			time.Sleep(time.Millisecond * 10)
		})
		_ = gc.Set(fmt.Sprintf("Test-%d", i), i)
	}
}
