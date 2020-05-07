package leetcode

func isIsomorphic2(s string, t string) bool {
    var res bool
    if len(s) == len(t) {
        hs:=makeHash(s)
        ht:=makeHash(t)
        res = true
        for i:=0;i<len(s);i++ {
            if hs[i]!=ht[i] {
                res = false
                break
            }
        }
    }
    return res
}

func makeHash(s string) []int {
    var res []int
    hash := make(map[byte]int)
    for i:=0;i<len(s);i++ {
        _,ok := hash[s[i]]
        if !ok {
            hash[s[i]] = i
        }
        res = append(res,hash[s[i]])
    }
    return res
}

func isIsomorphic(s string, t string) bool {
    var res bool
    if len(s) == len(t) {
        res = true
        hs := make(map[byte]int)
        ht := make(map[byte]int)
        for i:=0;i<len(s);i++ {
            if _,ok := hs[s[i]]; !ok {
                hs[s[i]] = i
            }
            if _,ok := ht[t[i]]; !ok {
                ht[t[i]] = i
            }
            if hs[s[i]] != ht[t[i]] {
                res = false
                break
            }
        }
    }
    return res
}
