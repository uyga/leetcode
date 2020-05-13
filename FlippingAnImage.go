package leetcode

func flipAndInvertImage(A [][]int) [][]int {
    l:=len(A[0])/2
    for i:=0;i<len(A);i++ {
        for j:=0;j<l;j++ {
            A[i][j],A[i][len(A[0])-1-j] = A[i][len(A[0])-1-j],A[i][j]
            A[i][j] = 1 - A[i][j]
            A[i][len(A[0])-1-j] = 1 - A[i][len(A[0])-1-j]
        }
        if len(A[0])%2 != 0 {
            A[i][l] = 1 - A[i][l]
        }
    }
    return A
}
