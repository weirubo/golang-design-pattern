package v1

import "fmt"

type singletonV1 struct {
}

var instance *singletonV1

func init() {
	instance = new(singletonV1)
	fmt.Printf("%p\n", instance)
}

func GetInstance() *singletonV1 {
	return instance
}
