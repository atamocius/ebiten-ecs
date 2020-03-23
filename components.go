package main

type velocity struct {
	VX float32
	VY float32
}

type position struct {
	X float32
	Y float32
}

type shape struct {
	Primitive primitive
}

type renderable struct {
	Flag bool
}
