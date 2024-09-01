func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	firstBuy, firstSell := -prices[0], 0
	secondBuy, secondSell := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		firstBuy = max(firstBuy, -prices[i])
		firstSell = max(firstSell, firstBuy+prices[i])
		secondBuy = max(secondBuy, firstSell-prices[i])
		secondSell = max(secondSell, secondBuy+prices[i])
	}

	return secondSell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
