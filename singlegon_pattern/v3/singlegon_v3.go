package v3

import (
	"fmt"
	"sync"
)

type singlegonV3 struct {
}

var instance *singlegonV3

var mu sync.Mutex

func GetInstance() *singlegonV3 {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		instance = new(singlegonV3)
	}
	fmt.Printf("%p\n", instance)
	return instance
}