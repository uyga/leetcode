package leetcode

func minCostClimbingStairs(cost []int) int {
    climbingStairsCache = make(map[int]int)
    return min(climbingStairsRecurcive(len(cost)-1,cost),climbingStairsRecurcive(len(cost)-2,cost))
}

func climbingStairsRecurcive(n int, cost []int) int {
    if n < 2 {
        return cost[n]
    }
    if _,ok := climbingStairsCache[n-1]; !ok {
        climbingStairsCache[n-1] = climbingStairsRecurcive(n-1,cost)
    }
    if _,ok := climbingStairsCache[n-2]; !ok {
        climbingStairsCache[n-2] = climbingStairsRecurcive(n-2,cost)
    }
    return cost[n] + min(climbingStairsCache[n-1],climbingStairsCache[n-2])
}

func min(n,m int) int {
    if n<m {
        return n
    } else {
        return m
    }
}
