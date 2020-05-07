package leetcode

func hammingWeight(num uint32) int {
    var res int
    for i:=0;i<32;i++ {
        if num & (1 << i) != 0 {
            res++
        }
    }
    return res
}
