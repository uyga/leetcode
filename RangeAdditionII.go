package leetcode

func maxCount(m int, n int, ops [][]int) int {
    var minx, miny int = m,n
    for _,op := range ops {
        if op[0] < minx {
            minx = op[0]
        }
        if op[1] < miny {
            miny = op[1]
        }
    }
    return minx*miny
}
