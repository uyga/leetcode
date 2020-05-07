package leetcode

func getSum(a int, b int) int {
    for b!=0 {
        var c int = a&b
        a=a^b
        b=c<<1
    }
    return a
}
