package leetcode

func checkRecord(s string) bool {
    var res bool = true
    if len(s) != 0 {
        var as,ls int
        var i,l int = 0, len(s)-1
        for i <= l {
            if s[i] == 'A' {
                as++
                if as > 1 {
                    res = false
                    i = l+1
                }
                i++
            } else if s[i] == 'L' {
                for i <= l && s[i] == 'L' {
                    ls++
                    if ls > 2 {
                        res = false
                        i = l+1
                    }
                    i++
                }
            } else {
                i++
            }
            ls = 0
        }
    }
    return res
}
