package main

import "golang-design-pattern/02_factory_pattern/abstract_factory"

func main() {
	// 简单工厂
	//simpleFactory := new(simple_factory.SimpleFactory)
	//cocaCola := simpleFactory.Produce("CocaCola")
	//cocaCola.Kind()
	//cocaCola.Name()
	//sprite := simpleFactory.Produce("Sprite")
	//sprite.Kind()
	//sprite.Name()

	// 工厂方法
	//cocaColaFactory := new(factory_method.CocaColaFactory)
	//cocaCola := cocaColaFactory.Produce()
	//cocaCola.Kind()
	//cocaCola.Name()
	//spriteFactory := new(factory_method.SpriteFactory)
	//sprite := spriteFactory.Produce()
	//sprite.Kind()
	//sprite.Name()

	// 抽象工厂
	cocaFactory := new(abstract_factory.CocaFactory)
	cocaCola := cocaFactory.ProduceCola()
	cocaCola.ColaKind()
	cocaCola.ColaName()
	cocaSprite := cocaFactory.ProduceSprite()
	cocaSprite.SpriteKind()
	cocaSprite.SpriteName()
	pepsiFactory := new(abstract_factory.PepsiFactory)
	pepsiCola := pepsiFactory.ProduceCola()
	pepsiCola.ColaKind()
	pepsiCola.ColaName()
	pepsiSprite := pepsiFactory.ProduceSprite()
	pepsiSprite.SpriteKind()
	pepsiSprite.SpriteName()
}
