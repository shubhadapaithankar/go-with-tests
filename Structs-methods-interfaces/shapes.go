package shapes

import "math"

// Shape interface will allow for polymorphism for different shapes.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle represents the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle represents a circle by its radius.
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the circumference of the circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle represents the dimensions of a triangle.
type Triangle struct {
	Base   float64
	Height float64
	SideA  float64
	SideB  float64
	SideC  float64
}

// Area calculates the area of the triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Perimeter calculates the perimeter of the triangle.
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}
