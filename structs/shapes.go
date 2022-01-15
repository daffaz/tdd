package main

const PI = 3.14159265358979323846264338327950288419716939937510582097494459

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return PI * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
