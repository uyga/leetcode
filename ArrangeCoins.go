package leetcode

func arrangeCoins(n int) int {
    var res int
    total := n
    for i := 1; i <= total; i ++ {
        n -= i
        if n >= 0 {
            res ++
        } else {
            break
        }
    }
    return res
}
