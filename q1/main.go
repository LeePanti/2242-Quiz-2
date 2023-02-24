package main

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
