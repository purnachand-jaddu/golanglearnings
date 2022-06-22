package main

import (
	"fmt"
)

type Playable struct {
}

type Bat struct {
	Playable

	Name string
}

func (Playable) play() {
	fmt.Println("Playing")
}

func (b *Bat) getName() string {
	return b.Name
}

func main() {
	bat := Bat{Name: "Kookaburra"}
	bat.play()
	fmt.Println(bat.getName())
}
