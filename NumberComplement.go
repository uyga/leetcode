package leetcode

import "math"

func findComplement(num int) int {
    if num == 0 {
        return 1
    }
    var res, i int
    for num > 0 {
        res += int(math.Pow(2.0,float64(i))) * (1^(num%2))
        i++
        num/=2
    }
    return res
}
