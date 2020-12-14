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
 * Same departure time:
 *
 * If all busses arrive at the same time,
 * we don't need to care about the order anymore.
 *
 * This approach stores the input-position.
 *
 * Currently this solution lacks in having negative
 * input when the input<position.
 */

var (
	mtx       sync.Mutex
	count     = 0
	doneCount = 0
)

func main() {

	lines := util.LoadString("input2")
	ts := make([]int, 0)
	i := 0
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

		ts = append(ts, num-i)
		i++
	}

	maxWorkers := runtime.GOMAXPROCS(0)
	fmt.Println("Max workers:", maxWorkers)
	sem := semaphore.NewWeighted(int64(maxWorkers))
	step := uint64(100000000)
	ch := make(chan uint64, 0)

	result := uint64(0)
	for t := uint64(0); result == 0; t += step {

		if sem.TryAcquire(1) {
			go worker(t, t+step, ts, sem, ch)
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
	start, end uint64,
	ts []int,
	sem *semaphore.Weighted,
	resultCh chan uint64) {

	defer sem.Release(1)
	defer func() {
		mtx.Lock()
		count--
		doneCount++
		mtx.Unlock()
	}()

	for t := start; t < end; t++ {
		if checkTs(t, ts) {
			resultCh <- t
			return
		}
	}
}

func checkTs(startTime uint64, ts []int) bool {
	for _, t := range ts {
		if startTime%uint64(t) != 0 {
			return false
		}
	}

	return true
}

func unique(in []int) []int {
	seen := make(map[int]bool, 0)
	result := make([]int, 0)

	for num := range in {

		if _, ok := seen[num]; !ok {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}
