package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
	"time"

	ch "computational_geometry_2/convex_hull"
)

func generatePointsInSquare(n int, sideLength float64) []ch.Point {
	points := make([]ch.Point, n)

	for i := 0; i < n; i++ {
		x := rand.Float64() * sideLength
		y := rand.Float64() * sideLength
		points[i] = ch.Point{X: x, Y: y}
	}

	return points
}

func generatePointsInCircle(n int, radius float64) []ch.Point {
	points := make([]ch.Point, n)

	for i := 0; i < n; i++ {
		r := radius * math.Sqrt(rand.Float64())
		theta := rand.Float64() * 2 * math.Pi
		x := r * math.Cos(theta)
		y := r * math.Sin(theta)
		points[i] = ch.Point{X: x, Y: y}
	}

	return points
}

func generatePointsOnCurve(n int, xBound float64) []ch.Point {
	points := make([]ch.Point, n)

	for i := 0; i < n; i++ {
		x := (rand.Float64() * 2 * xBound) - xBound
		y := -(x * x)
		points[i] = ch.Point{X: x, Y: y}
	}

	return points
}

func writeBenchmarkLineToCSV(writer *csv.Writer, n, sortTime int, runtimes, programCounters, backtrackCounters, bridgeCounters []int) {

	// Write the results as a new row
	record := []string{strconv.Itoa(n), strconv.Itoa(sortTime)}
	for i := 0; i < 4; i++ {
		record = append(record, strconv.Itoa(runtimes[i]))
	}
	for i := 0; i < 4; i++ {
		record = append(record, strconv.Itoa(programCounters[i]))
	}
	for i := 0; i < 4; i++ {
		record = append(record, strconv.Itoa(backtrackCounters[i]))
	}
	for i := 0; i < 4; i++ {
		record = append(record, strconv.Itoa(bridgeCounters[i]))
	}

	err := writer.Write(record)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	writer.Flush()
}

func main() {
	sizes := []int{65536000}
	// for i := 0; i < 9; i++ {
	// 	sizes = append(sizes, sizes[len(sizes)-1]*2)
	// }
	iterations := 10
	floatConstant := 10000.0

	for testClass := 2; testClass < 3; testClass++ {

		var fileName string
		if testClass == 0 {
			fileName = "square"
		} else if testClass == 1 {
			fileName = "circle"
		} else {
			fileName = "curve"
		}

		fmt.Println("Starting ", fileName, " benchmarks.")

		file, err := os.OpenFile("data/"+fileName+"/"+fileName+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		writer := csv.NewWriter(file)
		err = writer.Write([]string{"n", "sortTime", "seqRuntime", "p1Runtime", "p3Runtime", "p6Runtime", "seqCounter", "p1Counter", "p3Counter", "p6Counter", "bckCounterSeq", "bckCounterP1", "bckCounterP3", "bckCounterP6", "bridgeCounterSeq", "bridgeCounterP1", "bridgeCounterP3", "bridgeCounterP6"})
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		for _, n := range sizes {
			for i := 4; i < iterations; i++ {
				fmt.Printf("Class:\t%6s | Size:\t%9d | Iteration\t%d\n", fileName, n, i+1)

				var points []ch.Point
				if testClass == 0 {
					points = generatePointsInSquare(n, floatConstant)
				} else if testClass == 1 {
					points = generatePointsInCircle(n, floatConstant)
				} else if testClass == 2 {
					points = generatePointsOnCurve(n, floatConstant)
				} else {
					panic("WRONG CLASS COUNT")
				}

				fmt.Println("Generated Points")
				startSort := time.Now()
				ch.SortPointsByX(points)
				elapsedSort := time.Since(startSort).Nanoseconds()
				fmt.Println("Sorted Points")

				times := []int{}
				programCounters := []int{}
				backtrackCounters := []int{}
				bridgeCounters := []int{}

				fmt.Print("Started Sequential Graham Scan")
				startSeq := time.Now()
				_, progCtrSeq, bckCtrSeq := ch.INC_CH(points)
				elapsedSeq := time.Since(startSeq).Nanoseconds()
				fmt.Println("\rFinished Sequential Graham Scan")

				fmt.Print("Started Parallel Graham Scan (p=1)")
				startP1 := time.Now()
				_, progCtrP1, bckCtrP1, bridgeCtrP1 := ch.PAR_GS(points, 1)
				elapsedP1 := time.Since(startP1).Nanoseconds()
				fmt.Println("\rFinished Parallel Graham Scan (p=1)")

				fmt.Print("Started Parallel Graham Scan (p=3)")
				startP3 := time.Now()
				_, progCtrP3, bckCtrP3, bridgeCtrP3 := ch.PAR_GS(points, 3)
				elapsedP3 := time.Since(startP3).Nanoseconds()
				fmt.Println("\rFinished Parallel Graham Scan (p=3)")

				fmt.Print("Started Parallel Graham Scan (p=6)")
				startP6 := time.Now()
				_, progCtrP6, bckCtrP6, bridgeCtrP6 := ch.PAR_GS(points, 6)
				elapsedP6 := time.Since(startP6).Nanoseconds()
				fmt.Println("\rFinished Parallel Graham Scan (p=6)")

				times = append(times, int(elapsedSeq), int(elapsedP1), int(elapsedP3), int(elapsedP6))
				programCounters = append(programCounters, progCtrSeq, progCtrP1, progCtrP3, progCtrP6)
				backtrackCounters = append(backtrackCounters, bckCtrSeq, bckCtrP1, bckCtrP3, bckCtrP6)
				bridgeCounters = append(bridgeCounters, 0, bridgeCtrP1, bridgeCtrP3, bridgeCtrP6)

				fmt.Println("Writing results...")
				writeBenchmarkLineToCSV(writer, n, int(elapsedSort), times, programCounters, backtrackCounters, bridgeCounters)
				fmt.Println()
			}
		}

		file.Close()
	}

}
