package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func generatePointsInSquare(n int, sideLength float64) []Point {
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		x := rand.Float64() * sideLength
		y := rand.Float64() * sideLength
		points[i] = Point{x, y}
	}

	return points
}

func generatePointsInCircle(n int, radius float64) []Point {
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		r := radius * math.Sqrt(rand.Float64())
		theta := rand.Float64() * 2 * math.Pi
		x := r * math.Cos(theta)
		y := r * math.Sin(theta)
		points[i] = Point{x, y}
	}

	return points
}

func generatePointsOnCurve(n int, xBound float64) []Point {
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		x := (rand.Float64() * 2 * xBound) - xBound
		y := -(x * x)
		points[i] = Point{x, y}
	}

	return points
}

func main() {
	// points := generatePointsOnCurve(500, 100)
	points := []Point{{0, 3}, {2, 2}, {1, 1}, {2, 1}, {3, 0}, {0, 0}, {3, 3}}

	fmt.Println(PAR_GS(points, 2))
	// hull := INC_CH(points)

	// fmt.Println("Convex Hull")
	// fmt.Printf("points = [")
	// for idx, p := range points {
	// 	fmt.Printf("(%f, %f)", p.x, p.y)
	// 	if idx < len(points)-1 {
	// 		fmt.Printf(", ")
	// 	} else {
	// 		fmt.Printf("]\n")
	// 	}
	// }

	// fmt.Printf("hull_points = [")
	// for idx, p := range hull {
	// 	fmt.Printf("(%f, %f)", p.x, p.y)
	// 	if idx < len(hull)-1 {
	// 		fmt.Printf(", ")
	// 	} else {
	// 		fmt.Printf("]\n")
	// 	}
	// }

}
