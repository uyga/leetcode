package leetcode

func isPalindrome(s string) bool {
    if len(s) == 0 {
        return true
    }
    var i,j int
    j = len(s)
    for i=0;i<len(s);i++ {
        j--
        for i<len(s) && isNotAlphaNum(s[i]) {
            i++
        }
        for j>=0 && isNotAlphaNum(s[j]) {
            j--
        }
        if j < i {
            break
        }
        if i < len(s) && j >=0 && !equal(s[i],s[j]) {
            return false
        }
    }
    return true
}

func equal(a,b byte) bool {
    if a >= 'A' && b >= 'A' && a < b {
        return a == b - 32
    } else if a >= 'A' && b >= 'A' && a > b {
        return a == b + 32
    }
    return a == b
}

func isNotAlphaNum(char byte) bool {
    return char < '0' || char > '9' && char < 'A' || char > 'Z' && char < 'a' || char > 'z'
}
