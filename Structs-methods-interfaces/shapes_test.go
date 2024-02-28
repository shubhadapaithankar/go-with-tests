package shapes

import (
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 12, Height: 6}, 72.0},
		{Circle{Radius: 10}, 314.1592653589793},
		{Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 10, Height: 10}, 40.0},
		{Circle{Radius: 10}, 62.83185307179586},
		{Triangle{SideA: 3, SideB: 4, SideC: 5}, 12.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
