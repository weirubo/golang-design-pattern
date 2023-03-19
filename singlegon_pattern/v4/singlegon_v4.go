package v4

import (
	"fmt"
	"sync"
)

type singletonV4 struct {
}

var instance *singletonV4
var mu sync.Mutex

func GetInstance() *singletonV4 {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = new(singletonV4)
		}
	}
	fmt.Printf("%p\n", instance)
	return instance
}
