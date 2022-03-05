package main

import (
	"fmt"
	"math"
)

// Cuboid

type Cuboid struct {
	width  float64
	length float64
	height float64
}

func (c Cuboid) Volume() float64 {
	return c.length * c.height * c.width
}

func (c Cuboid) SurfaceArea() float64 {
	return 2 * (c.length*c.width + c.length*c.height + c.height*c.width)
}

// Sphere

type Sphere struct {
	radius float64
}

func (s Sphere) Volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(s.radius, 3)
}

func (s Sphere) SurfaceArea() float64 {
	return 2 * math.Pi * math.Pow(s.radius, 2)
}

type Solid interface {
	Volume() float64
	SurfaceArea() float64
}

func PrintInfo(i Solid) {
	fmt.Println("Volume of solid", i.Volume())
	fmt.Println("SurfaceArea of solid", i.SurfaceArea())
}

// func main() {
// 	var c Cuboid = Cuboid{2.0, 3.0, 4.0}
// 	PrintInfo(c)

// 	var s Sphere = Sphere{10.0}
// 	PrintInfo(s)

// }
