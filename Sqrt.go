package leetcode

import "math"

func mySqrt(x int) int {
    var res int = x
    if x > 1 {
        res = x / (2 * int(math.Log10(float64(x)) * 10))
        for x <= res * res {
            res = res / (2 * int(math.Log10(float64(res)) * 10))
        }
        for x > res * res {
            res++
        }
        if res*res > x {
            res--
        }
    }
    return res
}
