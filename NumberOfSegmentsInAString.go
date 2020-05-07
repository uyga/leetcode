package leetcode

func countSegments(s string) int {
    var res int
    if len(s) > 0 {
        i := 0
        l := len(s)-1
        for i<=l {
            for i<=l && isSpace(s[i]) {
                i++
            }
            if i<=l {
                res++
            }
            for i<=l && !isSpace(s[i]) {
                i++
            }
        }
    }
    return res
}

func isSpace(c byte) bool {
    return c == 9 || c == 32
}
