package twentytwo

type ProfitHistory [4]int

func TotalProfit(profits []map[ProfitHistory]int, seq ProfitHistory) int {
	total := 0

	for _, profit := range profits {
		if val, ok := profit[seq]; ok {
			total += val
		}
	}

	return total
}

func (s SecretNumber) ProfitDevelopment() map[ProfitHistory]int {
	prices := []int{s.Price()}
	history := make(map[ProfitHistory]int)

	for i := range 2000 {
		s = s.Next()
		prices = append(prices, s.Price())

		if len(prices) >= 5 {
			// Calculate profit history for the past four iterations
			changes := [...]int{
				prices[i-2] - prices[i-3],
				prices[i-1] - prices[i-2],
				prices[i] - prices[i-1],
				prices[i+1] - prices[i],
			}
			if _, exists := history[changes]; !exists {
				history[changes] = s.Price()
			}
		}
	}

	return history
}
