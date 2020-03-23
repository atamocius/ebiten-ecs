package main

type primitive int

const (
	box primitive = iota
	circle
)

func (p primitive) String() string {
	return [...]string{
		"box",
		"circle",
	}[p]
}
