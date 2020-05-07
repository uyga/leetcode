package leetcode

import "math"

func fib3(N int) int {
    if N <= 0 {
        return 0
    } else if N == 1 {
        return 1
    } else {
        return fib(N-1)+fib(N-2)
    }
}

var hash map[int]int

func fib2(N int) int {
    hash = make(map[int]int, N+1)
    return rec(N)
}

func rec(N int) int {
    if N <= 0 {
        return 0
    } else if N == 1 {
        return 1
    } else {
        if _,ok := hash[N-1]; !ok {
            hash[N-1] = rec(N-1)
        }
        if _,ok := hash[N-2]; !ok {
            hash[N-2] = rec(N-2)
        }
        return hash[N-1]+hash[N-2]
    }
}

func fib(N int) int {
    var goldenRatio float64 = (1.0 + math.Sqrt(5.0))/2.0
    return int(math.Round(math.Pow(goldenRatio, float64(N)) / math.Sqrt(5.0)))
}
