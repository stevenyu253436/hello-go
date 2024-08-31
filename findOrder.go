func findOrder(numCourses int, prerequisites [][]int) []int {
	// 構建圖和入度數組
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, prereq := range prerequisites {
		graph[prereq[1]] = append(graph[prereq[1]], prereq[0])
		inDegree[prereq[0]]++
	}

	// 使用隊列進行拓扑排序
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := []int{}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		order = append(order, course)

		for _, nextCourse := range graph[course] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	// 如果排序的課程數等於總課程數，返回排序結果
	if len(order) == numCourses {
		return order
	}

	// 否則返回空數組，表示無法完成所有課程
	return []int{}
}

func main() {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	order := findOrder(numCourses, prerequisites)
	fmt.Println(order) // [0, 1, 2, 3] 或 [0, 2, 1, 3]
}
