package main

import . "fmt"

type mover interface {
	move()
}

type eater interface {
	eat()
}

type slipper interface {
	slip()
}

type animal struct {
	name       string
	movingType string
	feedType1  string
	feedType2  string
}

func (a *animal) move(name string, movingType string) {
	Printf("%s %s", a.name, a.movingType)
}

func (a *animal) eat1(name string) {
	Printf("%s eat %s", a.name, a.feedType1)
}

func (a *animal) eat2(name string) {
	Printf("%s eat %s", a.feedType2)
}

func (a *animal) slip(name string) {
	Printf("%s slipping", a.name)
}

func main() {
	cow := animal{"Cow", "top-top", "grass", "silage"}
	horse := animal{"Horse", "gallop", "oats", "hay"}
	dog := animal{"Dog", "run", "bone", "meat"}

}
