package simple_factory

import "fmt"

// IDrink 抽象产品 - 饮料
type IDrink interface {
	Kind() // 抽象方法 - 类别
	Name() // 抽象方法 - 名称
}

// CocaCola 具体产品 - 可口可乐
type CocaCola struct {
}

// Kind 具体方法
func (c *CocaCola) Kind() {
	fmt.Println("carbonated drinks")
}

// Name 具体方法
func (c *CocaCola) Name() {
	fmt.Println("CocaCola")
}

// Sprite 具体产品 - 雪碧
type Sprite struct {
}

// Kind 具体方法
func (s *Sprite) Kind() {
	fmt.Println("carbonated drinks")
}

// Name 具体方法
func (s *Sprite) Name() {
	fmt.Println("Sprite")
}

// SimpleFactory 工厂
type SimpleFactory struct {
}

// Produce 生产 - 返回值（抽象产品）
func (s *SimpleFactory) Produce(name string) (drink IDrink) {
	if name == "CocaCola" {
		drink = new(CocaCola)
	} else if name == "Sprite" {
		drink = new(Sprite)
	}
	return
}
