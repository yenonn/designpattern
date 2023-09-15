package patterns

import "fmt"

type Object interface {
	Area() float64
	Height() float64
}

type VolumeCalculator struct {
	object Object
}

func (calculator VolumeCalculator) Volume() float64 {
	return calculator.object.Area() * calculator.object.Height()
}

// Implementing object: triangular
type TriangularObject struct {
	base   float64
	length float64
	height float64
}

func (t TriangularObject) Area() float64 {
	return 0.5 * t.base * t.length
}

func (t TriangularObject) Height() float64 {
	return t.height
}

// Implementing object: rectanglar
type RectangularObject struct {
	length float64
	width  float64
	height float64
}

func (r RectangularObject) Area() float64 {
	return r.length * r.width
}

func (r RectangularObject) Height() float64 { return r.height }

func Bridge() {
	triangular := TriangularObject{base: 0.4, length: 0.6, height: 4.6}
	triVolumeCalculator := VolumeCalculator{object: triangular}
	fmt.Println(triVolumeCalculator.Volume())
	rectangular := RectangularObject{length: 0.4, width: 0.2, height: 0.8}
	recVolumeCalculator := VolumeCalculator{object: rectangular}
	fmt.Println(recVolumeCalculator.Volume())
}
