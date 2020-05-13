package leetcode

import "math"

func maxAbsoluteDifference(A []int )  (int) {
    max1 := -2147483648.0
    min1 := 2147483647.0
    max2 := -2147483648.0
    min2 := 2147483647.0
    for i:=0;i<len(A);i++ {
        max1 = math.Max(max1, float64(A[i] + i))
        min1 = math.Min(min1, float64(A[i] + i))
        max2 = math.Max(max2, float64(A[i] - i))
        min2 = math.Min(min2, float64(A[i] - i))
    }
    return int(math.Max(max1 - min1, max2 - min2))
}
