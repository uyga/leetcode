package leetcode

func isPowerOfThree(n int) bool {
    if n <= 1 {
        return false
    }
    for n > 1 {
        if n % 3 == 0 {
            n /= 3
        } else {
            return false
        }
    }
    return true
}
