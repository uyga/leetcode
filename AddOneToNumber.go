package leetcode

func addOneToNumber(A []int )  ([]int) {
    A[len(A)-1]++
    i:=0
    for A[i] == 0 {
        i++
    }
    A = A[i:]
    i=len(A)-1
    for i>0 && A[i] > 9 {
        A[i-1] = A[i-1] + A[i]/10
        A[i] = A[i]%10
        i--
    }
    if A[0] > 9 {
        A = append(A, A[0]/10)
        A[0] = A[0] % 10
        A = append(A[len(A)-1:],A[:len(A)-1]...)
    }
    return A
}
