package leetcode

//O(N)
func isReverse(s string, t string) bool {
    var res bool = false
    var cnt int = 0
    if len(s) == len(t) {
        for i := 0; i < len(s); i++ {
            if s[i] == t[len(s)-1-i] {
                cnt++
            }
        }
        if len(s) == cnt {
            res = true
        }
    }
    return res
}
