package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
)

func generatePointsInSquare(n int, sideLength float64) []Point {
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		x := float64(int(rand.Float64() * sideLength))
		y := float64(int(rand.Float64() * sideLength))
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

// SavePointsToCSV saves points and hull points to a CSV file
func SavePointsToCSV(points []Point, hull []Point, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	writer.Write([]string{"Main Points", "Convex Hull Points"})

	// Find the max length between points and hull
	maxLen := len(points)
	if len(hull) > maxLen {
		maxLen = len(hull)
	}

	// Write points
	for i := 0; i < maxLen; i++ {
		var mainPoint, hullPoint string
		if i < len(points) {
			mainPoint = strconv.FormatFloat(points[i].x, 'f', 2, 64) + "," + strconv.FormatFloat(points[i].y, 'f', 2, 64)
		}
		if i < len(hull) {
			hullPoint = strconv.FormatFloat(hull[i].x, 'f', 2, 64) + "," + strconv.FormatFloat(hull[i].y, 'f', 2, 64)
		}
		writer.Write([]string{mainPoint, hullPoint})
	}

	return nil
}

// generateRandomPoints
func generateRandomPoints(n int, boundX, boundY float64) []Point {
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		x := (rand.Float64() * 2 * boundX) - boundX // Random x within [-boundX, boundX]
		y := (rand.Float64() * 2 * boundY) - boundY // Random y within [-boundY, boundY]
		points[i] = Point{x, y}
	}

	return points
}

func main() {
	// points := generateRandomPoints(150, 100, 100) // Generate 150 random points in a 200x200 box

	// points := generatePointsInSquare(15, 1000)
	// points := []Point{{64, 792}, {91, 9}, {111, 348}, {128, 3}, {162, 623}, {327, 794}, {436, 597}, {488, 221}, {562, 129}, {588, 398}, {728, 182}, {748, 927}, {913, 176}, {984, 716}, {990, 865}}
	points := []Point{{0, 3}, {2, 2}, {1, 1}, {2, 1} /* */, {3, 0}, {0, 0}, {3, 3}, {4, 2}}

	res1 := INC_CH(points, 0)
	res2 := INC_CH(points, 2)

	fmt.Printf("POINTS\n")
	for _, p := range points {
		fmt.Printf("(%.0f, %.0f), ", p.x, p.y)
	}

	fmt.Printf("\nSEQ CH\n")
	for _, p := range res1 {
		fmt.Printf("(%.0f, %.0f), ", p.x, p.y)
	}
	fmt.Printf("\nPAR CH\n")
	for _, p := range res2 {
		fmt.Printf("(%.0f, %.0f), ", p.x, p.y)
	}

	fmt.Println()
}
