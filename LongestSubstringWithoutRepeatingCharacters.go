package leetcode

func lengthOfLongestSubstring2(s string) int {
    var res int
    if len(s) > 0 {
        i := 1
        l := len(s)-1
        res = 1
        hash := make(map[byte]bool)
        hash[s[0]]=true
        cnt := 1
        prev := i
        for i <= l {
            if _,ok := hash[s[i]]; !ok {
                cnt++
                hash[s[i]]=true
            } else {
                if cnt > res {
                    res = cnt
                }
                hash = make(map[byte]bool)
                hash[s[prev]]=true
                cnt = 1
                i = prev
                prev++
            }
            i++
        }
        if cnt > res {
            res = cnt
        }
    }
    return res
}

func lengthOfLongestSubstring(s string) int {
    var res int
    if len(s) > 0 {
        start := 0
        hash := make(map[byte]int)
        for i:=0;i<len(s);i++ {
            pos, ok := hash[s[i]]
            hash[s[i]] = i
            if ok && pos >= start {
                start = pos + 1
            } else if i-start+1 > res {
                res = i-start+1
            }
        }
    }
    return res
}
