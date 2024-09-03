package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	events := []Event{}
	for _, b := range buildings {
		events = append(events, Event{b[0], -b[2]}) // start event
		events = append(events, Event{b[1], b[2]})  // end event
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].x != events[j].x {
			return events[i].x < events[j].x
		}
		return events[i].h < events[j].h
	})

	heap := &MaxHeap{}
	heap.Push(0) // Ground level
	prevMax := 0
	var result [][]int

	for _, event := range events {
		if event.h < 0 {
			heap.Push(-event.h)
		} else {
			heap.Remove(event.h)
		}

		currMax := heap.Max()
		if currMax != prevMax {
			result = append(result, []int{event.x, currMax})
			prevMax = currMax
		}
	}

	return result
}

type Event struct {
	x int
	h int
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
	heap.Fix(h, len(*h)-1)
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Remove(val int) {
	for i, v := range *h {
		if v == val {
			heap.Remove(h, i)
			break
		}
	}
}

func (h *MaxHeap) Max() int {
	if len(*h) > 0 {
		return (*h)[0]
	}
	return 0
}

func main() {
	buildings := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}
	result := getSkyline(buildings)
	fmt.Println(result)
}
