package v3

import (
	"fmt"
	"sync"
)

type singletonV3 struct {
}

var instance *singletonV3

var mu sync.Mutex

func GetInstance() *singletonV3 {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		instance = new(singletonV3)
	}
	fmt.Printf("%p\n", instance)
	return instance
}
