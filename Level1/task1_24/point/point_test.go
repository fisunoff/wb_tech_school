package point

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name     string
		p1       *Point
		p2       *Point
		expected float64
	}{
		{
			name:     "origin to (3,4)",
			p1:       NewPoint(0, 0),
			p2:       NewPoint(3, 4),
			expected: 5.0,
		},
		{
			name:     "same point",
			p1:       NewPoint(1, 1),
			p2:       NewPoint(1, 1),
			expected: 0.0,
		},
		{
			name:     "negative coordinates",
			p1:       NewPoint(-1, -1),
			p2:       NewPoint(2, 3),
			expected: 5.0,
		},
		{
			name:     "decimal values",
			p1:       NewPoint(1.5, -2.3),
			p2:       NewPoint(-3.7, 4.1),
			expected: math.Sqrt((1.5+3.7)*(1.5+3.7) + (-2.3-4.1)*(-2.3-4.1)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p1.Distance(tt.p2)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("Distance() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestNewPoint(t *testing.T) {
	p := NewPoint(3.14, 2.71)
	if p == nil {
		t.Fatal("NewPoint returned nil")
	}
	zero := NewPoint(0, 0)
	d := p.Distance(zero)
	expected := math.Sqrt(3.14*3.14 + 2.71*2.71)
	if math.Abs(d-expected) > 1e-9 {
		t.Errorf("NewPoint created wrong point: distance from origin is %v, expected %v", d, expected)
	}
}
