package leetcode

var climbingStairsCache map[int]int

func climbStairs(n int) int {
    if n <= 3 {
        return n
    }
    if climbingStairsCache == nil {
        climbingStairsCache = make(map[int]int)
    }
    if _,ok:=climbingStairsCache[n-1];!ok {
        climbingStairsCache[n-1] = climbStairs(n-1)
    }
    if _,ok:=climbingStairsCache[n-2];!ok {
        climbingStairsCache[n-2] = climbStairs(n-2)
    }
    return climbingStairsCache[n-1]+climbingStairsCache[n-2]
}
