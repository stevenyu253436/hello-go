func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	// 初始化布尔数组，全部设为true
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	// 使用埃拉托斯特尼筛法
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	// 统计质数的个数
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}
