package point

import (
	"fmt"
	"math"
)

// Point представляет точку на плоскости с приватными координатами
type Point struct {
	x, y float64
}

// NewPoint — конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance вычисляет расстояние до другой точки
func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

// String строковое представление координат
func (p *Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.x, p.y)
}
