package leetcode

import "math"

func coverPoints(A []int , B []int )  (int) {
    cnt := 0
    for i:=0;i<len(A)-1;i++ {
        cnt += path(A[i],A[i+1],B[i],B[i+1])
    }
    return cnt
}

func path(x1,x2,y1,y2 int) int {
    dx := math.Abs(float64(x2-x1))
    dy := math.Abs(float64(y2-y1))
    return int(math.Max(dx,dy))
}
