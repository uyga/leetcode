package leetcode

func isHappy(n int) bool {
    hash := make(map[int]bool)
    for n != 1 {
        n = getPowSum(n)
        _, val := hash[n]
        if val {
            return false
        }
        hash[n] = true
    }
    return true
}

func getPowSum(n int) int {
    var sum int
    for n>0 {
        sum = sum + (n%10) * (n%10)
        n=n/10
    }
    return sum
}
