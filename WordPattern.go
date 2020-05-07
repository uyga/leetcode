package leetcode

func wordPattern(pattern string, str string) bool {
    wmask := wordMask(pattern)
    smask := sentenceMask(str)
    if len(wmask) != len(smask) {
        return false
    }
    for i:=0;i<len(wmask);i++ {
        if wmask[i] != smask[i] {
            return false
        }
    }
    return true
}

func wordMask(s string) []int {
    var res []int
    hash := make(map[byte]int)
    for i:=0;i<len(s);i++ {
        if _,ok := hash[s[i]]; !ok {
            hash[s[i]] = i
        }
        res = append(res, hash[s[i]])
    }
    return res
}

func sentenceMask(s string) []int {
    var res []int
    hash := make(map[string]int)
    var word string
    var cnt int
    for i:=0;i<=len(s);i++ {
        if i == len(s) || s[i] == ' ' {
            if _,ok := hash[word]; !ok {
                hash[word] = cnt
            }
            res = append(res, hash[word])
            word = ""
            cnt++
        } else {
            word += string(s[i])
        }
    }
    return res
}
