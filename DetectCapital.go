package leetcode

func detectCapitalUse(word string) bool {
    if len(word) == 0 {
        return true
    }
    var res bool = true
    var ethalon bool = isCapital(word[0])
    for i:=1;i<len(word);i++ {
        if isCapital(word[i]) != ethalon {
            if i == 1 && ethalon == true {
                ethalon = false
            } else {
                res = false
                break
            }
        }
    }
    return res
}

func isCapital(c byte) bool {
    return c >= 'A' && c <= 'Z'
}
