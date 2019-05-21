package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Height: 12, Width: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
		}
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	}

// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		want := 100.0
// 		checkArea(t, rectangle, want)
// 	})
// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		want := 314.1592653589793
// 		checkArea(t, circle, want)
// 	})
// }
