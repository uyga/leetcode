package leetcode

func maxProfitII(prices []int) int {
    var profit,transaction int = 0,0
    if len(prices) > 0 {
        var minprice int = prices[0]
        for i:=1;i<len(prices);i++ {
            if minprice > prices[i] {
                //buy
                minprice = prices[i]
            } else if transaction < prices[i] - minprice {
                //sell
                transaction = prices[i] - minprice
                profit += transaction
                transaction = 0
                minprice = prices[i]
            }
        }
    }
    return profit
}
