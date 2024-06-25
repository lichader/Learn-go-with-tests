package structs

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Height: 6, Width: 12}, want: 72.0}, // This is a valid test case
		{shape: Circle{10}, want: math.Pi * 100},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
