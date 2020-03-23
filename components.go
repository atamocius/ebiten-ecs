package main

type velocity struct {
	vX float32
	vY float32
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
