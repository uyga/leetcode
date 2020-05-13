package leetcode

func transposeMatrix(A [][]int) [][]int {
    var res [][]int
    if len(A) > 0 {
        i,j := 0, 0
        li := len(A)-1
        lj := len(A[i])-1
        var tmp []int
        for j<=lj {
            i=0
            tmp=[]int{}
            for i<=li {
                tmp = append(tmp, A[i][j])
                i++
            }
            res = append(res, tmp)
            j++
        }
    }
    return res
}
