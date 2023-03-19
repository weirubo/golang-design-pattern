package v2

import "fmt"

type singlegonV2 struct {
}

var instance *singlegonV2

func GetInstance() *singlegonV2 {
	if instance == nil {
		instance = new(singlegonV2)
	}
	fmt.Printf("%p\n", instance)
	return instance
}
