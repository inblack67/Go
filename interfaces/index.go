package main

import (
	"fmt"
	"math"
)

type util interface{
	area() float64
}

type circle struct{
	radius float64
}

type rectangle struct{
	width, height float64
}

func (c circle) area() float64{
	area := math.Pi * c.radius * c.radius
	return math.Ceil(area)
}

func (r rectangle) area() float64{
	area := r.width * r.height
	return math.Ceil(area)
}

func getArea(s util) float64{
	return s.area()
}

func main(){
	c1 := circle{radius: 5}
	rect1 := rectangle{width: 12, height: 12}
	fmt.Println("circle area", getArea(c1))
	fmt.Println("rectangle area", getArea(rect1))
}
