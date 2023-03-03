package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) GeQuater() int {
	// need to implement
	return 0
}

type Line struct {
	start, finish Point
}

func (l *Line) GetLenght() float64 {
	return math.Sqrt(math.Pow((l.finish.x-l.finish.x), 2) + math.Pow((l.finish.y-l.finish.y), 2))
}

type Square struct {
	side Line
}

func (s *Square) GetPerimeter() float64 {
	return 4 * s.side.GetLenght()
}

func (s *Square) GetArea() float64 {
	return math.Pow(s.side.GetLenght(), 2)
}

func main() {
	square := Square{
		side: Line{
			start: Point{
				x: 4,
				y: 4,
			},
			finish: Point{
				x: 0,
				y: 4,
			},
		},
	}
	fmt.Println(square)
	fmt.Println(square.GetPerimeter())
	fmt.Println(square.GetArea())
}
