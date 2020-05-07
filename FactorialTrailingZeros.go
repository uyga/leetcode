package leetcode

func trailingZeroes(n int) int {
    var res int
    for n >= 5 {
        n/=5
        res+=n
    }
    return res
}
