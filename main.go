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

	sortPointsX(points)

	fmt.Println("Starting test")
	times := []int{}
	programCounters := []int{}

	start1 := time.Now()
	_, progCtr1 := GS(points)
	elapsed1 := time.Since(start1).Nanoseconds()

	start2 := time.Now()
	_, progCtr2 := PAR_GS(points, 1)
	elapsed2 := time.Since(start2).Nanoseconds()

	start3 := time.Now()
	_, progCtr3 := PAR_GS(points, 3)
	elapsed3 := time.Since(start3).Nanoseconds()

	start4 := time.Now()
	_, progCtr4 := PAR_GS(points, 6)
	elapsed4 := time.Since(start4).Nanoseconds()

	times = append(times, int(elapsed1), int(elapsed2), int(elapsed3), int(elapsed4))
	programCounters = append(programCounters, progCtr1, progCtr2, progCtr3, progCtr4)

	fmt.Println(times)
	fmt.Println(programCounters)
}
