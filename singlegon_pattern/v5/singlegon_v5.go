package v5

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singlegonV5 struct {
}

var instance *singlegonV5

var mu sync.Mutex

var done uint32

func GetInstance() *singlegonV5 {
	if atomic.LoadUint32(&done) == 0 {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			defer atomic.StoreUint32(&done, 1)
			instance = new(singlegonV5)
		}
	}
	fmt.Printf("%p\n", instance)
	return instance
}
