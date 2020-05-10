package leetcode

func spiralOrder(matrix [][]int) []int {
    var res []int
    if len(matrix) > 0 {
        sw := 0
        w := len(matrix[sw])-1
        sh := 0
        h := len(matrix)-1
        for sw<=w && sh<=h {
            //top row
            for i:=sw;i<=w;i++ {
                res = append(res, matrix[sw][i])
            }
            sh++
            //right column
            for i:=sh;i<=h;i++ {
                res = append(res, matrix[i][w])
            }
            w--
            //bottom row
            for i:=w;i>=sw && h >= sh;i-- {
                res = append(res, matrix[h][i])
            }
            h--
            //left column
            for i:=h;i>=sh && w >= sw;i-- {
                res = append(res, matrix[i][sw])
            }
            sw++
        }
    }
    return res
}
