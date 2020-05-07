package leetcode

func getRow(rowIndex int) []int {
    var res []int
    var cnt int
    for i:=0;i<=rowIndex;i++ {
        cnt++
        var tmp []int
        for t:=0;t<cnt;t++ {
            if i < 2 || (t == 0 || t == cnt-1) {
                tmp = append(tmp, 1)
            } else {
                tmp = append(tmp, res[t] + res[t-1])
            }
        }
        res = tmp
    }
    return res
}
