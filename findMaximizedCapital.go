import (
	"container/heap"
	"sort"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := make([][2]int, len(profits))
	for i := range profits {
		projects[i] = [2]int{capital[i], profits[i]}
	}

	// 按照需要的資本從小到大排序專案
	sort.Slice(projects, func(i, j int) bool {
		return projects[i][0] < projects[j][0]
	})

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	index := 0
	for i := 0; i < k; i++ {
		// 將當前資金下可啟動的專案加入最大堆
		for index < len(projects) && projects[index][0] <= w {
			heap.Push(maxHeap, projects[index][1])
			index++
		}

		// 如果最大堆中有專案，選擇利潤最大的專案
		if maxHeap.Len() > 0 {
			w += heap.Pop(maxHeap).(int)
		} else {
			// 如果沒有可選擇的專案，提前結束
			break
		}
	}

	return w
}
