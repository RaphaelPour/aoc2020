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

var (
	mtx       sync.Mutex
	count     = 0
	doneCount = 0
)

func main() {

	lines := util.LoadString("input")
	ts := make([]uint64, 0)
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

		ts = append(ts, uint64(num-i))

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
	ts []uint64,
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

func checkTs(startTime uint64, ts []uint64) bool {
	for _, t := range ts {
		if startTime%t != 0 {
			return false
		}
	}

	return true
}
