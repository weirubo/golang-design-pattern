package v4

import (
	"fmt"
	"sync"
)

type singlegonV4 struct {
}

var instance *singlegonV4
var mu sync.Mutex

func GetInstance() *singlegonV4 {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = new(singlegonV4)
		}
	}
	fmt.Printf("%p\n", instance)
	return instance
}
