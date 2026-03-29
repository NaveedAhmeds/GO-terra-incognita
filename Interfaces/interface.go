// Interface in go
// -> collection of method signatures
// -> every inheriting entity declares its own method but has the same signatures.
// -> interfaces are implemented implicitly

package main

// Example interface of a shape
type shape interface {
	area() float64
	perimiter() float64
}

// a random shape struct
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimiter() {
	return 2*r.width + 2*r.height
}
