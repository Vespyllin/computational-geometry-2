package main
// GiftWrapping finds the convex hull using the gift-wrapping algorithm (Jarvis March).
func GiftWrapping(points []Point) []Point {
    n := len(points)
    if n < 3 {
        return nil // Convex hull is not possible
    }

    hull := []Point{}

    // Find the leftmost point
    leftMost := 0
    for i := 1; i < n; i++ {
        if points[i].x < points[leftMost].x {
            leftMost = i
        }
    }

    p := leftMost
    for {
        hull = append(hull, points[p])
        q := (p + 1) % n

        for i := 0; i < n; i++ {
            if orientation(points[p], points[i], points[q]) == 2 {
                q = i
            }
        }

        p = q

        if p == leftMost {
            break
        }
    }

    return hull
}
