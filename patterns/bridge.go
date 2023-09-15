package patterns

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
	height float64
}

func (t TriangularObject) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t TriangularObject) Height() float64 {
	return t.height
}

func main() {
}
