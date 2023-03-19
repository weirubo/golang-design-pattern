package v6

import (
	"fmt"
	"sync"
)

type singlegonV6 struct {
}

var instance *singlegonV6

var once sync.Once

func GetInstance() *singlegonV6 {
	once.Do(func() {
		instance = new(singlegonV6)
	})
	fmt.Printf("%p\n", instance)
	return instance
}
