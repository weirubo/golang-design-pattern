package v2

import "fmt"

type singletonV2 struct {
}

var instance *singletonV2

func GetInstance() *singletonV2 {
	if instance == nil {
		instance = new(singletonV2)
	}
	fmt.Printf("%p\n", instance)
	return instance
}
