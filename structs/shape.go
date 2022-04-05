package main

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	w float64
	h float64
}

func (r Rectangle) Area() float64 {
	return r.w * r.h
}

type Circle struct {
	r float64
}

type Triangle struct {
	b float64
	h float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.w + rect.h)
}

func Area(rect Rectangle) float64 {
	return rect.w * rect.h
}

func (t Triangle) Area() float64 {
	return (t.b * t.h) * 0.5
}
