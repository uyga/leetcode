package leetcode

import "math"

func reverseInteger(x int) int {
    var res int = 0
    for i := x; i != 0; i = i / 10 {
        if res * 10 + i % 10 > math.MaxInt32 || res * 10 + i % 10 < math.MinInt32 {
            res = 0
            break
        }
        res = (res * 10) + i % 10
    }
    return res
}
