package main

import "sort"

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
	n := len(points)
	if n < 3 {
		return points
	}

	// Sort over x
	sort.Slice(points, func(i, j int) bool {
		if points[i].x == points[j].x {
			return points[i].y < points[j].y
		}
		return points[i].x < points[j].x
	})

	// Initialize upper hull
	upperHull := []Point{points[0], points[1]}

	// Consolidate points into the UH
	for i := 2; i < n; i++ {
		// While the accrued convex hull is not valid, remove points until it is valid again
		for len(upperHull) >= 2 && orientation(upperHull[len(upperHull)-2], upperHull[len(upperHull)-1], points[i]) == LEFT {
			upperHull = upperHull[:len(upperHull)-1]
		}

		// Otherwise append the point to the UH
		upperHull = append(upperHull, points[i])
	}

	// Initialize lower hull
	lowerHull := []Point{points[0], points[1]}

	// Consolidate points into the UH
	for i := 2; i < n; i++ {
		// While the accrued convex hull is not valid, remove points until it is valid again
		for len(lowerHull) >= 2 && orientation(lowerHull[len(lowerHull)-2], lowerHull[len(lowerHull)-1], points[i]) == RIGHT {
			lowerHull = lowerHull[:len(lowerHull)-1]
		}

		// Otherwise append the point to the UH
		lowerHull = append(lowerHull, points[i])
	}

	// Remove duplicate points
	upperHull = upperHull[:len(upperHull)-1]
	lowerHull = lowerHull[1:]
	reversePoints(lowerHull)
	return append(upperHull, lowerHull...)
}
