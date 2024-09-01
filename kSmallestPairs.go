import (
	"container/heap"
)

type Pair struct {
	sum  int
	i, j int
}

type MinHeap []Pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return nil
	}

	h := &MinHeap{}
	heap.Init(h)

	// Initialize the heap with the first element in nums1 paired with each element in nums2
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, Pair{nums1[i] + nums2[0], i, 0})
	}

	result := [][]int{}

	for k > 0 && h.Len() > 0 {
		pair := heap.Pop(h).(Pair)
		result = append(result, []int{nums1[pair.i], nums2[pair.j]})
		k--

		if pair.j+1 < len(nums2) {
			heap.Push(h, Pair{nums1[pair.i] + nums2[pair.j+1], pair.i, pair.j + 1})
		}
	}

	return result
}
