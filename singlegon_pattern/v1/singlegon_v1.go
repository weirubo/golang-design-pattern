package v1

import "fmt"

type singlegonV1 struct {
}

var instance *singlegonV1

func init() {
	instance = new(singlegonV1)
	fmt.Printf("%p\n", instance)
}

func GetInstance() *singlegonV1 {
	return instance
}
