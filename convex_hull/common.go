package convex_hull

import "sort"

type Point struct{ X, Y float64 }

type Line struct{ p1, p2 Point }

const (
	COLLINEAR = iota
	RIGHT
	LEFT
)

func orientation(p, q, r Point) int {
	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)

	if val == 0 {
		return COLLINEAR
	} else if val > 0 {
		return RIGHT
	} else {
		return LEFT
	}
}

func SortPointsByX(points []Point) {
	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})
}

// Used for graphing purposes
// func reversePoints(points []Point) {
// 	n := len(points)
// 	for i := 0; i < n/2; i++ {
// 		points[i], points[n-i-1] = points[n-i-1], points[i]
// 	}
// }
