package convex_hull

func GIFT_CH(points []Point) ([]Point, int) {
	counter := 0

	n := len(points)
	if n < 3 {
		return points, counter
	}

	hull := make([]Point, 0, len(points))

	leftMost := 0
	for i := 1; i < n; i++ {
		counter++
		if points[i].X < points[leftMost].X || (points[i].X == points[leftMost].X && points[i].Y < points[leftMost].Y) {
			leftMost = i
		}
	}

	p := leftMost
	for {
		counter++
		hull = append(hull, points[p])

		q := (p + 1) % n
		for i := 0; i < n; i++ {
			counter++
			if orientation(points[p], points[i], points[q]) == RIGHT {
				q = i
			}
		}

		p = q

		if p == leftMost {
			return hull, counter
		}
	}
}
