package leetcode

func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    }
    var bt,et int = 0, len(t)-1
    var bs,es int = 0, len(s)-1
    for bt <= et {
        if s[bs] == t[bt] {
            if bs >= es {
                return true
            }
            bs++
        }
        bt++
        if s[es] == t[et] {
            if bs >= es {
                return true
            }
            es--
        }
        et--
    }
    return false
}
