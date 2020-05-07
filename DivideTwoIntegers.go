package leetcode

import "math"

func divide2(dividend int, divisor int) int {
    if dividend == 0 {
        return 0
    }
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }
    var sign int = 1
    if dividend < 0 {
        sign *= -1
        dividend *= -1
    }
    if divisor < 0 {
        sign *= -1
        divisor *= -1
    }
    var cnt int
    for dividend > 0 {
        if dividend < divisor {
            break
        } else {
            dividend -= divisor
            cnt++
        }
    }
    return cnt*sign
}

func divide(dividend int, divisor int) int {
    var sign int = 1
    if dividend < 0 {
        sign *= -1
        dividend *= -1
    }
    if  divisor < 0 {
        sign *= -1
        divisor *= -1
    }
    var i,q,t int
    for i=31;i>=0;i-- {
        if t + (divisor << i) <= dividend {
            t += divisor << i
            q |= 1 << i

        }
    }
    if q*sign > math.MaxInt32 {
        return math.MaxInt32
    }
    return q*sign
}
