package leetcode

import "math"

func isPalindromeNumber(x int) bool {
    var res bool = true
    var cnt int
    var xlen int = int(math.Log10(float64(x)) + 1)
    for i := x; i != 0; i = i / 10 {
        cnt++
        if (i % 10) != (x / int(math.Pow(10,float64(xlen-cnt))) % 10) {
            res = false
            break
        }
    }
    return res
}
