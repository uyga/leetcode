package leetcode

func hammingDistance(x int, y int) int {
    var res int
    for i:=0;i<32;i++ {
        if (x>>i)&1 != (y>>i)&1 {
            res++
        }
    }
    return res
}
