package leetcode

func isUgly(num int) bool {
    if num <= 0 {
        return false
    }
    if num == 1 {
        return true
    }
    for num > 1 {
        cnt := 0
        if num%2 == 0 {
            num/=2
            cnt++
        }
        if num%3 == 0 {
            num/=3
            cnt++
        }
        if num%5 == 0 {
            num/=5
            cnt++
        }
        if cnt == 0 {
            return false
        }
    }
    return true
}
