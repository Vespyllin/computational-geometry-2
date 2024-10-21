package main

type Point struct{ x, y float64 }

type Line struct{ p1, p2 Point }

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

// Used for graphing purposes
// func reversePoints(points []Point) {
// 	n := len(points)
// 	for i := 0; i < n/2; i++ {
// 		points[i], points[n-i-1] = points[n-i-1], points[i]
// 	}
// }
