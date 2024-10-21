package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortPointsX(points []Point) {
	sort.Slice(points, func(i, j int) bool {
		if points[i].x == points[j].x {
			return points[i].y < points[j].y
		}
		return points[i].x < points[j].x
	})
}

// Graham Scan Starter
func INC_CH(points []Point, p int) []Point {
	// Sort by lowest x coordinate
	sortPointsX(points)

	var res []Point
	if p == 0 {
		res = GS(points)
	} else {
		res = PAR_GS(points, p)
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].x == res[j].x {
			return res[i].y < res[j].y
		}
		return res[i].x < res[j].x
	})

	return res

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
	avgSliceLen := len(points) / p
	remaining := len(points) % p

	var wg sync.WaitGroup
	wg.Add(p)

	hulls := make([][]Point, p)
	considered := 0
	for i := 0; i < p; i++ {
		sliceLen := avgSliceLen
		if i < remaining {
			sliceLen += 1
		}

		activeSlice := points[considered : considered+sliceLen]

		go func(pointsSlice []Point, rank int) {
			defer wg.Done()
			hulls[i] = GS(pointsSlice)
		}(activeSlice, i)

		considered += sliceLen
	}

	wg.Wait()

	// for i := 0; i < p-1; i++ {
	// 	x := []int{}
	// 	for j := i + 1; j < p; j++ {
	// 		x := append(x, j) // append Tangent between Ui and Uj
	// 	}

	// 	y := min(x)          // Find tangent with smallest rotation
	// 	i := indexOf(min(x)) // skip ahead to the hull targeted by that tangent

	// }

	finalHull := hulls[0]
	// for i := 1; i < len(hulls); i++ {
	// 	// finalHull = mergeHulls(finalHull, hulls[i])
	// 	// sortPointsX(finalHull)
	// }
	fmt.Println(finalHull)

	return finalHull
}

func mergeHulls(leftHull, rightHull []Point) []Point {
	// Find top and bottom tangent lines between the convex hulls
	bottom := getTangentialPoints(leftHull, rightHull, false)
	top := getTangentialPoints(leftHull, rightHull, true)

	// fmt.Println("top", top)
	// fmt.Println("bottom", bottom)
	// Filter out the points now inside the merged hull
	mergedHull := []Point{}
	// First points (bottom[0] and top[0])
	if bottom[0].x < top[0].x || (bottom[0].x == top[0].x && bottom[0].y < top[0].y) {
		mergedHull = append(mergedHull, bottom[0], top[0])
	} else if bottom[0].x != top[0].x || bottom[0].y != top[0].y {
		mergedHull = append(mergedHull, top[0], bottom[0])
	} else {
		mergedHull = append(mergedHull, top[0])
	}

	// Second points (bottom[1] and top[1])
	if bottom[1].x < top[1].x || (bottom[1].x == top[1].x && bottom[1].y < top[1].y) {
		mergedHull = append(mergedHull, bottom[1], top[1])
	} else if bottom[1].x != top[1].x || bottom[1].y != top[1].y {
		mergedHull = append(mergedHull, top[1], bottom[1])
	} else {
		mergedHull = append(mergedHull, top[1])
	}

	// Right hand side of OR case is to keep those points collinear with the bridge boundaries
	if !(bottom[0].x == top[0].x && bottom[0].y == top[0].y) {
		for _, p := range leftHull {
			if orientation(bottom[0], top[0], p) == LEFT || (!(p.x == bottom[0].x && p.y == bottom[0].y) && !(p.x == top[0].x && p.y == top[0].y) && orientation(bottom[0], p, top[0]) == COLLINEAR) {
				mergedHull = append(mergedHull, p)
			}
		}
	}

	if !(bottom[1].x == top[1].x && bottom[1].y == top[1].y) {
		for _, p := range rightHull {
			if orientation(bottom[1], top[1], p) == RIGHT || (!(p.x == bottom[1].x && p.y == bottom[1].y) && !(p.x == top[1].x && p.y == top[1].y) && orientation(bottom[1], p, top[1]) == COLLINEAR) {
				mergedHull = append(mergedHull, p)
			}
		}
	}

	return mergedHull
}

// Returns left and right hull points forming a bridge
func getTangentialPoints(leftHull, rightHull []Point, top bool) []Point {
	var TARGET_ORIENTATION int
	if top {
		TARGET_ORIENTATION = LEFT
	} else {
		TARGET_ORIENTATION = RIGHT
	}

	leftIdx := len(leftHull) - 1
	rightIdx := 0

	if !top {
		fmt.Println("START")
		fmt.Println(leftHull, rightHull)
		fmt.Println(leftHull[leftIdx], rightHull[rightIdx])
	}

	for {

		if !top {
			fmt.Println(leftIdx, rightIdx)
		}
		prevLeftIdx := (leftIdx + 1) % len(leftHull)
		nextLeftIdx := (leftIdx - 1 + len(leftHull)) % len(leftHull)
		prevRightIdx := (rightIdx - 1 + len(rightHull)) % len(rightHull)
		nextRightIdx := (rightIdx + 1) % len(rightHull)

		if orientation(leftHull[leftIdx], rightHull[rightIdx], rightHull[nextRightIdx]) == TARGET_ORIENTATION { // Attempt to adjust right bridge point to the right
			rightIdx = nextRightIdx
		} else if orientation(leftHull[leftIdx], rightHull[rightIdx], rightHull[prevRightIdx]) == TARGET_ORIENTATION { // Attempt to adjust right bridge point to the left
			rightIdx = prevRightIdx
		} else if orientation(leftHull[nextLeftIdx], leftHull[leftIdx], rightHull[rightIdx]) == TARGET_ORIENTATION { // Attempt to adjust left bridge point to the left
			leftIdx = nextLeftIdx
		} else if orientation(leftHull[prevLeftIdx], leftHull[leftIdx], rightHull[rightIdx]) == TARGET_ORIENTATION { // Attempt to adjust left bridge point to the right
			leftIdx = prevLeftIdx
		} else {
			return []Point{leftHull[leftIdx], rightHull[rightIdx]}
		}

		if !top {
			fmt.Println(leftHull[leftIdx], rightHull[rightIdx])
		}
	}

}
