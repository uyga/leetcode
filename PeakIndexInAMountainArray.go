package leetcode

func peakIndexInMountainArray(A []int) int {
    max := 0
    for i:=0;i<len(A);i++ {
        if A[max] < A[i] {
            max = i
        }
    }
    return max
}
