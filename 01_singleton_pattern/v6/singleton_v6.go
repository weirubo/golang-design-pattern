package v6

import (
	"fmt"
	"sync"
)

type singletonV6 struct {
}

var instance *singletonV6

var once sync.Once

func GetInstance() *singletonV6 {
	once.Do(func() {
		instance = new(singletonV6)
	})
	fmt.Printf("%p\n", instance)
	return instance
}
