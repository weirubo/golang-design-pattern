package factory_method

import "fmt"

// IDrink 抽象产品
type IDrink interface {
	Kind() // 抽象方法
	Name() // 抽象方法
}

// CocaCola 具体产品
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

// Sprite 具体产品
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

// IFactory 抽象工厂
type IFactory interface {
	Produce() IDrink // 抽象方法
}

// CocaColaFactory 具体工厂
type CocaColaFactory struct {
}

// Produce 具体方法
func (c *CocaColaFactory) Produce() (drink IDrink) {
	drink = new(CocaCola)
	return
}

// SpriteFactory 具体工厂
type SpriteFactory struct {
}

// Produce 具体方法
func (s *SpriteFactory) Produce() (drink IDrink) {
	drink = new(Sprite)
	return
}
