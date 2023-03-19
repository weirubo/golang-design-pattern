package v5

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singletonV5 struct {
}

var instance *singletonV5

var mu sync.Mutex

var done uint32

func GetInstance() *singletonV5 {
	if atomic.LoadUint32(&done) == 0 {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			defer atomic.StoreUint32(&done, 1)
			instance = new(singletonV5)
		}
	}
	fmt.Printf("%p\n", instance)
	return instance
}
