package leetcode

func backspaceCompare(S string, T string) bool {
    var res bool
    s:=buildStack(S)
    t:=buildStack(T)
    if len(s) == len(t) {
        res = true
        for i:=0;i<len(s);i++ {
            if s[i] != t[i] {
                res = false
                break
            }
        }
    }
    return res
}

func buildStack(s string) string {
    var stack []byte
    for i:=0;i<len(s);i++ {
        if s[i] != '#' {
            stack = append(stack, s[i])
        } else {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        }
    }
    return string(stack)
}
