package main

type velocity struct {
	VX, VY float32
}

type position struct {
	X, Y float32
}

type shape struct {
	Primitive primitive
}

type renderable struct {
	Flag bool
}
