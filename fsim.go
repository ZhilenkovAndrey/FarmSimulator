package main

import (
	. "fmt"
	"math/rand"
	"time"
)

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

func (a animal) String() string {
	return a.name
}

func (a *animal) move() {
	Printf("%s %s\n", a.name, a.movingType)
}

func (a *animal) eat() {
	Printf("%s eat %s\n", a.name, a.feedChoise())
}

func (a *animal) slip() {
	Printf("%s slipping\n", a.name)
}

func (a *animal) feedChoise() string {
	var s string
	b := rand.Intn(2)
	switch b {
	case 0:
		s = Sprintf("%s", a.feedType1)
	case 1:
		s = Sprintf("%s", a.feedType2)
	}
	return s
}

func life(an []animal) {
	for i := 0; i < 24; i++ {

		Println()
		Printf(" TIME: %d:00\n", i)

		if 0 <= i && i < 6 {
			for j := range an {
				an[j].slip()
			}
		}
		if 6 <= i && i < 12 {
			for j := range an {
				an[j].eat()
			}
		}
		if 12 <= i && i < 18 {
			for j := range an {
				an[j].move()
			}
		}
		if 18 <= i && i < 24 {
			for j := range an {
				an[j].eat()
			}
		}

		time.Sleep(time.Second)
	}
}

func animalAppearence() []animal {
	cow := animal{"Cow", "top-top", "grass", "silage"}
	horse := animal{"Horse", "gallop", "oats", "hay"}
	dog := animal{"Dog", "run", "bone", "meat"}

	return []animal{cow, horse, dog}
}

func main() {
	life(animalAppearence())
}
