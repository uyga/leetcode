package leetcode

func isPerfectSquare(num int) bool {
    if num <= 0 {
        return false
    }
    if num == 1 {
        return true
    }
    low := 2
    high := num
    for low < high {
        mid := (high - low)/2 + low
        if mid * mid < num {
            low = mid+1
        } else if mid * mid > num {
            high = mid
        } else {
            return true
        }
    }
    return false
}
