package main

import (
	"sort"
	"sync"
)

// Used for graphing purposes
func reversePoints(points []Point) {
	n := len(points)
	for i := 0; i < n/2; i++ {
		points[i], points[n-i-1] = points[n-i-1], points[i]
	}
}

type Point struct{ x, y float64 }

const (
	COLLINEAR = iota
	RIGHT
	LEFT
)

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

// Graham Scan
func INC_CH(points []Point) []Point {
	// Sort by lowest x coordinate
	sort.Slice(points, func(i, j int) bool {
		if points[i].x == points[j].x {
			return points[i].y < points[j].y
		}
		return points[i].x < points[j].x
	})

	return GS(points)
}

func GS(points []Point) []Point {
	n := len(points)
	if n < 3 {
		return points
	}

	upperHull := []Point{points[0], points[1]}
	for i := 2; i < n; i++ {
		// If the next point makes the convex hull invalid, pop points off the hull until it doesn't
		for len(upperHull) >= 2 && orientation(upperHull[len(upperHull)-2], upperHull[len(upperHull)-1], points[i]) == LEFT {
			upperHull = upperHull[:len(upperHull)-1]
		}

		// Append the point to the hull
		upperHull = append(upperHull, points[i])
	}

	// Repeat for the lower hull (Other direction)
	lowerHull := []Point{points[0], points[1]}
	for i := 2; i < n; i++ {
		for len(lowerHull) >= 2 && orientation(lowerHull[len(lowerHull)-2], lowerHull[len(lowerHull)-1], points[i]) == RIGHT {
			lowerHull = lowerHull[:len(lowerHull)-1]
		}

		lowerHull = append(lowerHull, points[i])
	}

	// Remove duplicate points
	upperHull = upperHull[:len(upperHull)-1]
	lowerHull = lowerHull[1:]
	return append(upperHull, lowerHull...)
}

func PAR_GS(points []Point, p int) []Point {
	// Sort over x
	sort.Slice(points, func(i, j int) bool {
		if points[i].x == points[j].x {
			return points[i].y < points[j].y
		}
		return points[i].x < points[j].x
	})

	avgSliceLen := len(points) / p
	remaining := len(points) % p

	var wg sync.WaitGroup
	wg.Add(p)

	CHs := make([][]Point, p)
	considered := 0
	for i := 0; i < p; i++ {
		sliceLen := avgSliceLen
		if i < remaining {
			sliceLen += 1
		}

		activeSlice := points[considered : considered+sliceLen]

		go func(pointsSlice []Point, rank int) {
			defer wg.Done()
			CHs[i] = GS(pointsSlice)
		}(activeSlice, i)

		considered += sliceLen
	}

	wg.Wait()

	// TODO: Merge CHs

	return points
}
