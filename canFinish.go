func canFinish(numCourses int, prerequisites [][]int) bool {
	// 建立每門課程的相鄰課程列表
	graph := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		graph[prereq[1]] = append(graph[prereq[1]], prereq[0])
	}

	// 記錄每個節點的訪問狀態：0 = 未訪問, 1 = 訪問中, 2 = 已訪問
	visitStatus := make([]int, numCourses)

	// DFS 檢查是否存在環
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if visitStatus[course] == 1 { // 檢測到環
			return false
		}
		if visitStatus[course] == 2 { // 已經檢查過，沒有環
			return true
		}

		visitStatus[course] = 1 // 設置為訪問中
		for _, nextCourse := range graph[course] {
			if !dfs(nextCourse) {
				return false
			}
		}
		visitStatus[course] = 2 // 設置為已訪問

		return true
	}

	// 遍歷每個課程進行 DFS 檢查
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}, {0, 1}}
	result := canFinish(numCourses, prerequisites)
	fmt.Println(result) // false
}
