package main

import (
	"encoding/csv"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
)

type Point struct{ x, y float64 }

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

// moved from Graham
const (
	COLLINEAR = iota
	RIGHT
	LEFT
)

func generatePointsOnCurve(n int, xBound float64) []Point {
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		x := (rand.Float64() * 2 * xBound) - xBound
		y := -(x * x)
		points[i] = Point{x, y}
	}

	return points
}

// moved from Graham
func orientation(p, q, r Point) int {
	val := (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)
	if val == 0 {
		return COLLINEAR
	} else if val > 0 {
		return RIGHT
	} else {
		return LEFT
	}
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
        x := (rand.Float64() * 2 * boundX) - boundX  // Random x within [-boundX, boundX]
        y := (rand.Float64() * 2 * boundY) - boundY  // Random y within [-boundY, boundY]
        points[i] = Point{x, y}
    }

    return points
}


func main() {
	points := generateRandomPoints(150, 100, 100)  // Generate 150 random points in a 200x200 box

	// points := generatePointsOnCurve(500, 100)
	// points := []Point{{0, 3}, {2, 2}, {1, 1}, {2, 1}, {3, 0}, {0, 0}, {3, 3}}

	// fmt.Println(PAR_GS(points, 2))
	//
	// fmt.Println(GiftWrapping(points))

	hull := GiftWrapping(points)
	hullParGs := PAR_GS(points, 2)
	hullParGsP := PAR_GS1(points, 2)
	hullGs := INC_CH(points)

	SavePointsToCSV(points, hull, "GiftWrapping.csv")
	SavePointsToCSV(points, hullGs, "GrahamIncScan.csv")
	SavePointsToCSV(points, hullParGs, "ParallelGrahamScanPaul.csv")
	SavePointsToCSV(points, hullParGsP, "ParallelGrahamScanMe.csv")

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
