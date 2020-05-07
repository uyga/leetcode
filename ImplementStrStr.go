package leetcode

func strStr(haystack string, needle string) int {
    var res, cnt int = -1, 0
    if needle == "" {
        res = 0
    } else {
        for i:=0;i<len(haystack);i++ {
            if len(needle) + i > len(haystack) {
                break
            }
            if haystack[i] == needle[0] {
                cnt = 0
                for j:=0;j<len(needle);j++ {
                    if haystack[i+j] == needle[j] {
                        cnt++
                    }
                }
                if cnt == len(needle) {
                    res = i
                    break
                }
            }
        }
    }
    return res
}
