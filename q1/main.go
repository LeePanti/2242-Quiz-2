package main

import "fmt"

type instrument interface {
	sound() string
	playingTechnique() string
}

type guitar struct {
	material string
	types    string
}

func (g guitar) sound() string {
	return "strum strum"
}

func (g guitar) playingTechnique() string {
	return "plucking"
}

type drum struct {
	material string
	types    string
}

func (d drum) sound() string {
	return "boom"
}

func (d drum) playingTechnique() string {
	return "striking"
}

type cricket struct {
	color string
	size  int
}

func (c cricket) sound() string {
	return "chirp"
}

func main() {
	myGuitar := guitar{
		material: "wood",
		types:    "acoustic",
	}
	myDrum := drum{
		material: "wood",
		types:    "Bass",
	}

	smallCricket := cricket{
		color: "green",
		size:  3,
	}

	printInfo(myGuitar)
	printInfo(myDrum)
	//printInfo(smallCricket)

	fmt.Println(smallCricket.sound())
}

func printInfo(i instrument) {
	fmt.Println("This Instrument has a", i.sound(), "sound and is played by", i.playingTechnique(), "it.")
}
