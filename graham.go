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

	sortPointsX(res)

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

	return upperHull
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

	bridges := []Line{}
	for i := 0; i < p-1; i++ {
		fmt.Println(i)
		tanLines := []Line{}
		for j := i + 1; j < p; j++ {
			fmt.Println("j", j)
			tanLines = append(tanLines, getTangentialPoints(hulls[i], hulls[j])) // append Tangent between Ui and Uj
		}

		minTan, minTanIdx := findMinRotationTangent(tanLines)
		bridges = append(bridges, minTan)

		i = i + minTanIdx
		fmt.Println("new i", i)
	}

	fmt.Println("bridges")
	fmt.Println(bridges)

	finalHull := hulls[0]
	// for i := 1; i < len(hulls); i++ {
	// 	// finalHull = mergeHulls(finalHull, hulls[i])
	// 	// sortPointsX(finalHull)
	// }
	// fmt.Println(finalHull)

	return finalHull
}

func findMinRotationTangent(lines []Line) (Line, int) {
	minTangent := lines[0]
	minIdx := 0

	for idx, line := range lines[1:] {
		// Compare rotation by checking the orientation
		if orientation(line.p1, line.p2, minTangent.p2) == LEFT {
			minTangent = line
			minIdx = idx
		}
	}

	return minTangent, minIdx
}

// Returns left and right hull points forming a bridge
func getTangentialPoints(leftHull, rightHull []Point) Line {
	TARGET_ORIENTATION := LEFT

	leftIdx := len(leftHull) - 1
	rightIdx := 0

	for {
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
			return Line{leftHull[leftIdx], rightHull[rightIdx]}
		}
	}

}
