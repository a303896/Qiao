package designPattern

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Fruit interface {
	Eat()
}

type Apple struct {
	Name string
}

type Orange struct {
	Name string
}

type FruitFactory struct{}

type Cherry struct {
	Name string
}

func (p *Apple) Eat() {
	fmt.Printf("eat apple %s\n", p.Name)
}

func (p *Orange) Eat() {
	fmt.Printf("eat orange %s\n", p.Name)
}

func (p *Cherry) Eat() {
	fmt.Printf("eat cherry %s\n", p.Name)
}

func NewApple(name string) Fruit {
	return &Apple{name}
}

func NewOrange(name string) Fruit {
	return &Orange{name}
}

func NewCherry(name string) Fruit {
	return &Cherry{name}
}

func NewFruitFactory() *FruitFactory {
	return &FruitFactory{}
}

func (f FruitFactory) CreateFruit(t string) (Fruit, error) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	name := strconv.Itoa(r.Int())
	switch t {
	case "Apple":
		return NewApple(name), nil
	case "Orange":
		return NewOrange(name), nil
	case "Cherry":
		return NewCherry(name), nil
	default:
		return nil, fmt.Errorf("wrong fruit type: %s", t)
	}
}

// 简单工厂方法改良
type creator func(name string) Fruit
type FruitFactory2 struct {
	creators map[string]creator
}

func NewFruitFactory2() *FruitFactory2 {
	return &FruitFactory2{
		map[string]creator{
			"apple":  NewApple,
			"orange": NewOrange,
			"cherry": NewCherry,
		},
	}
}

func (f *FruitFactory2) CreateFruit2(t string) (Fruit, error) {
	c, ok := f.creators[t]
	if !ok {
		return nil, fmt.Errorf("wrong fruit type: %s", t)
	}
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	name := strconv.Itoa(r.Int())
	return c(name), nil
}
