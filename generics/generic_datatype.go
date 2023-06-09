package main

import "fmt"

type Numbers interface {
	int | int16 | float32 | float64
}

type Strings interface {
	string
}

type GenericContainer[T Numbers | Strings] struct {
	values []T
	keys   []T
}

func (gc GenericContainer[T]) DisplayGenericContainer() {
	fmt.Println("Keys : ", gc.keys)
	fmt.Println("Values : ", gc.values)
}
