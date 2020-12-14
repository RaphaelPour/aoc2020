package main

import (
	"fmt"
	"github.com/RaphaelPour/aoc2020/util"
	"golang.org/x/sync/semaphore"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
 * Brute force:
 *
 * Check for each time step if the busses are
 * arriving in order.
 *
 * This version can check those steps in parallel.
 * The time can be incremented by the minimum
 * of the input.
 *
 * This solution takes a lot of time.
 */

var (
	mtx       sync.Mutex
	count     = 0
	doneCount = 0
)

type Departure struct {
	Index uint64
	BusID uint64
}

func main() {

	lines := util.LoadString("input")
	departures := make([]Departure, 0)
	i := 0
	max := 0
	min := 10000
	for _, cell := range strings.Split(lines[1], ",") {

		if cell == "x" {
			i++
			continue
		}

		num, err := strconv.Atoi(cell)
		if err != nil {
			fmt.Println("Cell is not a number", cell)
			return
		}
		departures = append(departures, Departure{
			Index: uint64(i),
			BusID: uint64(num),
		})
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
		i++
	}
	fmt.Println("min,max: ", min, max)

	maxWorkers := runtime.GOMAXPROCS(0)
	fmt.Println("Max workers:", maxWorkers)
	sem := semaphore.NewWeighted(int64(maxWorkers))
	step := uint64(1000000000)
	timeStep := uint64(min)
	ch := make(chan uint64, 0)

	result := uint64(0)
	for t := uint64(0); result == 0; t += step {

		if sem.TryAcquire(1) {
			go worker(t, t+step, timeStep, departures, sem, ch)
			mtx.Lock()
			count++
			mtx.Unlock()

			fmt.Printf(
				"\rRunning: %d Done: %d t: %d",
				count,
				doneCount,
				t,
			)
		}
		select {
		case result = <-ch:
		case <-time.After(100 * time.Millisecond):
		}
	}

	fmt.Println("\r>>", result, "<<")
	return

}

func worker(
	start, end, step uint64,
	deps []Departure,
	sem *semaphore.Weighted,
	resultCh chan uint64) {

	defer sem.Release(1)
	defer func() {
		mtx.Lock()
		count--
		doneCount++
		mtx.Unlock()
	}()

	for t := start; t < end; t += step {
		if checkDepartures(t, deps) {
			resultCh <- t
			return
		}
	}
}

func checkDepartures(startTime uint64, deps []Departure) bool {
	for i := 0; i < len(deps); i++ {
		t := deps[i].Index
		/*
			fmt.Printf("(%d+%d)%%%d=%d\n",
				t, startTime, deps[i].BusID,
				(t+startTime)%deps[i].BusID,
			)*/
		if (t+startTime)%deps[i].BusID != 0 {
			return false
		}
	}

	return true
}
