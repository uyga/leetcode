package leetcode

func generate(numRows int) [][]int {
    var res [][]int
    var cnt int = 0
    for i:=0;i<numRows;i++ {
        cnt++
        var tmp []int
        for t:=0;t<cnt;t++ {
            if i < 2 || (t == 0 || t == cnt-1) {
                tmp = append(tmp, 1)
            } else {
                tmp = append(tmp, res[i-1][t-1]+res[i-1][t])
            }
        }
        res = append(res, tmp)
    }
    return res
}
