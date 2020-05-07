package leetcode

//O(n^2)
func maxProfitBruteForce(prices []int) int {
    var profit int = 0
    for i := 0; i < len(prices); i++ {
        for j := len(prices)-1; j > i; j-- {
            if prices[j] - prices[i] > profit {
                profit = prices[j] - prices[i]
            }
        }
    }
    return profit
}

//O(n)
func maxProfit(prices []int) int {
    var profit int = 0
    if len(prices) > 0 {
        var minvalue int = prices[0]
        for i := 1; i < len(prices); i++ {
            if prices[i] < minvalue {
                minvalue = prices[i]
            } else if prices[i] - minvalue > profit {
                profit = prices[i] - minvalue
            }
        }
    }
    return profit
}
