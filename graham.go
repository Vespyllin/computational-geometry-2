package main

import (
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

func GS(points []Point) ([]Point, int) {
	counter := 0
	n := len(points)
	if n < 3 {
		return points, counter
	}

	upperHull := []Point{points[0], points[1]}
	for i := 2; i < n; i++ {
		// If the next point makes the convex hull invalid, pop points off the hull until it doesn't
		for len(upperHull) >= 2 && orientation(upperHull[len(upperHull)-2], upperHull[len(upperHull)-1], points[i]) == LEFT {
			upperHull = upperHull[:len(upperHull)-1]
			counter++
		}

		// Append the point to the hull
		upperHull = append(upperHull, points[i])
		counter++
	}

	return upperHull, counter
}

func PAR_GS(points []Point, p int) ([]Point, int) {
	counter := 0
	avgSliceLen := len(points) / p
	remaining := len(points) % p

	var wg sync.WaitGroup
	wg.Add(p)

	hulls := make([][]Point, p)
	parallelCounters := make([]int, p)

	considered := 0
	for i := 0; i < p; i++ {
		counter++
		sliceLen := avgSliceLen
		if i < remaining {
			sliceLen += 1
		}

		activeSlice := points[considered : considered+sliceLen]

		go func(pointsSlice []Point, rank int) {
			defer wg.Done()

			hull, subCounter := GS(pointsSlice)

			hulls[i] = hull
			parallelCounters[i] = subCounter
		}(activeSlice, i)

		considered += sliceLen
	}

	wg.Wait()

	maxParCtr := 0
	for _, parCtr := range parallelCounters {
		if parCtr > maxParCtr {
			maxParCtr = parCtr
		}
	}
	counter += maxParCtr

	finalHull := []Point{}
	lastL := 0
	lastR := 0
	for i := 0; i < p-1; {
		counter++
		tanLines := []Line{}
		tanLineHullIdx := [][]int{}
		for j := i + 1; j < p; j++ {
			counter++

			l, r, tanPointsCtr := getTangentialPoints(hulls[i], hulls[j])
			counter += tanPointsCtr

			tanLines = append(tanLines, Line{hulls[i][l], hulls[j][r]}) // Append tangent between Ui and Uj
			tanLineHullIdx = append(tanLineHullIdx, []int{l, r})        // Also keep track of the indexes for said tangent
		}

		minTanIdx, minTanRotCtr := findMinRotationTangentIdx(tanLines)
		counter += minTanRotCtr

		bridge := tanLines[minTanIdx]

		leftTangentPoint := tanLineHullIdx[minTanIdx][0]
		rightTangentPoint := tanLineHullIdx[minTanIdx][1]

		// If there are intermediary points append them to the hull before the bridge
		if lastL <= leftTangentPoint {
			intermediaryPoints := hulls[i][lastL:leftTangentPoint]
			finalHull = append(finalHull, intermediaryPoints...)
			finalHull = append(finalHull, bridge.p1, bridge.p2)
		} else {
			finalHull = append(finalHull, bridge.p2) // Otherwise append the next point on the bridge
		}

		lastL = rightTangentPoint + 1
		lastR = tanLineHullIdx[minTanIdx][1]

		i = i + 1 + minTanIdx
	}

	// Append any remaining points after the final bridge
	if p != 1 {
		lastR += 1
	}
	finalHull = append(finalHull, hulls[len(hulls)-1][lastR:]...)

	return finalHull, counter
}

func findMinRotationTangentIdx(lines []Line) (int, int) {
	counter := 0
	minIndex := 0

	for i := 1; i < len(lines); i++ {
		// Compare rotation by checking the orientation
		if orientation(lines[i].p1, lines[i].p2, lines[minIndex].p2) == RIGHT {
			minIndex = i
		}
		counter++
	}

	return minIndex, counter
}

// Returns left and right hull points forming a bridge
func getTangentialPoints(leftHull, rightHull []Point) (int, int, int) {
	counter := 0
	TARGET_ORIENTATION := LEFT

	leftIdx := len(leftHull) - 1
	rightIdx := 0

	for {
		counter++
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
			return leftIdx, rightIdx, counter
		}
	}
}
