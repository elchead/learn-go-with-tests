package main

func Area(width, height float64) float64 {
	return width * height
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
