package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
	"time"
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
	points := generatePointsInCircle(100000000, 1000)
	// points := []Point{{64, 792}, {91, 9}, {111, 348}, {128, 3}, {162, 623}, {327, 794}, {436, 597}, {488, 221}, {562, 129}, {588, 398}, {728, 182}, {748, 927}, {913, 176}, {984, 716}, {990, 865}}
	// points := []Point{{0, 3}, {2, 2}, {1, 1}, {2, 1}, {2, 2} /* */, {3, 0}, {0, 0}, {3, 3}, {4, 2}}
	// points := []Point{{0, 3}, {2, 6}, {4, 3}, {1, 5}, {2.5, 6} /* */, {4, 0}, {0, 0}, {3, 5}, {5, 1}}

	sortPointsX(points)

	start1 := time.Now()
	resS := INC_CH(points, 0)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	resP := INC_CH(points, 1)
	elapsed2 := time.Since(start2)

	fmt.Printf("%4d  (%10d)\n%4d  (%10d)\n\t===> %t\n", len(resS), elapsed1.Nanoseconds(), len(resP), elapsed2.Nanoseconds(), len(resS) == len(resP))
}
