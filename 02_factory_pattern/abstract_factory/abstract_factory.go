package abstract_factory

import "fmt"

// AbstractCola 抽象 Cola
type AbstractCola interface {
	ColaKind() // 抽象方法
	ColaName() // 抽象方法
}

// AbstractSprite 抽象 Sprite
type AbstractSprite interface {
	SpriteKind() // 抽象方法
	SpriteName() // 抽象方法
}

// AbstractFactory 抽象工厂
type AbstractFactory interface {
	ProduceCola() AbstractCola     // 抽象方法
	ProduceSprite() AbstractSprite // 抽象方法
}

// CocaBrandCola 可口品牌 具体 Cola 产品
type CocaBrandCola struct {
}

func (c *CocaBrandCola) ColaKind() {
	fmt.Println("Coca Brand carbonated drinks")
}

func (c *CocaBrandCola) ColaName() {
	fmt.Println("Coca Brand Cola")
}

// CocaBrandSprite 可口品牌 具体 Sprite 产品
type CocaBrandSprite struct {
}

func (c *CocaBrandSprite) SpriteKind() {
	fmt.Println("Coca Brand carbonated drinks")
}

func (c *CocaBrandSprite) SpriteName() {
	fmt.Println("Coca Brand Sprite")
}

// CocaFactory 可口品牌 具体工厂
type CocaFactory struct {
}

func (c *CocaFactory) ProduceCola() (cola AbstractCola) {
	cola = new(CocaBrandCola)
	return
}

func (c *CocaFactory) ProduceSprite() (sprite AbstractSprite) {
	sprite = new(CocaBrandSprite)
	return
}

// PepsiBrandCola 百事品牌 具体 Cola 产品
type PepsiBrandCola struct {
}

func (p *PepsiBrandCola) ColaKind() {
	fmt.Println("Pepsi Brand carbonated drinks")
}

func (p *PepsiBrandCola) ColaName() {
	fmt.Println("Pepsi Brand Cola")
}

// PepsiBrandSprite 百事品牌 具体 Sprite 产品
type PepsiBrandSprite struct {
}

func (p *PepsiBrandSprite) SpriteKind() {
	fmt.Println("Pepsi Brand carbonated drinks")
}

func (p *PepsiBrandSprite) SpriteName() {
	fmt.Println("Pepsi Brand Sprite")
}

// PepsiFactory 百事品牌 具体工厂
type PepsiFactory struct {
}

func (p *PepsiFactory) ProduceCola() (cola AbstractCola) {
	cola = new(PepsiBrandCola)
	return
}

func (p *PepsiFactory) ProduceSprite() (sprite AbstractSprite) {
	sprite = new(PepsiBrandSprite)
	return
}
