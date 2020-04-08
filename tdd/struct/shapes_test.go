package tdd

import (
	"math"
	"testing"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func Perimeter(reactangle Rectangle) float64 {
	return 2 * (reactangle.width + reactangle.height)
}

func Area(reactangle Rectangle) float64 {
	return reactangle.width * reactangle.height
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (t Triangle) Area() float64 {
	return (t.height * t.base) * 0.5
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }
	// t.Run("rectangles", func(t *testing.T) {

	// 	rectangle := Rectangle{12.0, 6.0}

	// 	got := Area(rectangle)
	// 	want := 72.0

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	circle := Circle{10}

	// 	got := Area(circle)
	// 	want := 314.1592653589793

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// })

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
