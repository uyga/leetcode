package leetcode

func titleToNumber(s string) int {
    var res int
    for i:=0;i<len(s);i++ {
        res = res * int('Z'-'A' + 1) + int(s[i] - 'A' + 1)
    }
    return res
}
