package convex_hull

func GIFT_CH(points []Point) ([]Point, int) {
	counter := 0

	n := len(points)
	if n < 3 {
		return points, counter
	}

	hull := []Point{}

	p := 0
	for {
		counter++
		hull = append(hull, points[p])

		q := (p + 1) % n
		for i := 0; i < n; i++ {
			if orientation(points[p], points[i], points[q]) == RIGHT {
				q = i
			}
		}

		p = q

		// Special termination condition for upper hull
		// Once we reach the rightmost point return the hull
		if points[q].X == points[len(points)-1].X && points[q].Y == points[len(points)-1].Y {
			hull = append(hull, points[p])
			return hull, counter
		}
	}
}
