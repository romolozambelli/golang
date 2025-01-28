package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Square", func(t *testing.T) {
		ret := Square{10, 12}
		expectedArea := float64(120)
		receivedArea := ret.Area()

		if expectedArea != receivedArea {

			t.Errorf("Received area %f is different of expected %f", receivedArea, expectedArea)
			//t.Fatalf("Received area %f is different of expected %f", receivedArea, expectedArea) - Stops the test
		}

	})

	t.Run("Circle", func(t *testing.T) {
		circ := Circle{10}
		expectedArea := float64(math.Pi * 100)
		receivedArea := circ.Area()

		if expectedArea != receivedArea {

			t.Errorf("Received area %f is different of expected %f", receivedArea, expectedArea)
			//t.Fatalf("Received area %f is different of expected %f", receivedArea, expectedArea) - Stops the test
		}
	})
}
