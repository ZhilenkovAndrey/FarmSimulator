package main

import . "fmt"

type mover interface {
	move()
}

type eater interface {
	eat()
}

type animal struct {
	name string
}

func (a *animal) move(name string) {
	Printf("%s moving", a.name)
}

func (a *animal) eat(name string) {
	Printf("%s eating", a.name)
}
