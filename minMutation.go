package main

import (
	"fmt"
)

func minMutation(startGene string, endGene string, bank []string) int {
	// 將基因庫轉換成集合，方便查詢
	bankSet := make(map[string]bool)
	for _, gene := range bank {
		bankSet[gene] = true
	}

	if _, found := bankSet[endGene]; !found {
		return -1
	}

	// 基因的四種可能變化
	genes := []rune{'A', 'C', 'G', 'T'}

	// 初始化隊列和訪問集合
	queue := []string{startGene}
	visited := make(map[string]bool)
	visited[startGene] = true

	mutations := 0

	for len(queue) > 0 {
		nextQueue := []string{}
		mutations++

		for _, currentGene := range queue {
			// 逐個字符嘗試突變
			for i := 0; i < len(currentGene); i++ {
				for _, gene := range genes {
					if rune(currentGene[i]) == gene {
						continue
					}

					mutatedGene := currentGene[:i] + string(gene) + currentGene[i+1:]

					if mutatedGene == endGene {
						return mutations
					}

					if _, found := bankSet[mutatedGene]; found && !visited[mutatedGene] {
						visited[mutatedGene] = true
						nextQueue = append(nextQueue, mutatedGene)
					}
				}
			}
		}

		queue = nextQueue
	}

	return -1
}

func main() {
	startGene := "AACCGGTT"
	endGene := "AAACGGTA"
	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	result := minMutation(startGene, endGene, bank)
	fmt.Println(result) // 輸出應為 2
}
