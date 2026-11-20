package util

import (
	"math"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	tests := []struct {
		name                               string
		latSrc, longSrc, latDest, longDest float64
		want                               float64
	}{
		{name: "one degree longitude", latSrc: 1, longSrc: 1, latDest: 1, longDest: 2, want: 111.178},
		{name: "zero coordinate", latSrc: 0, longSrc: 1, latDest: 1, longDest: 2, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateDistance(tt.latSrc, tt.longSrc, tt.latDest, tt.longDest)
			if math.Abs(got-tt.want) > 0.1 {
				t.Errorf("CalculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
