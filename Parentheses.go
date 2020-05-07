package leetcode

//O(n)
func isValid(s string) bool {
    var res bool = false
    if len(s) % 2 == 0 {
        var stack []byte
        hash := map[byte]byte{')' : '(', '}' : '{', ']' : '['}
        for i:=0; i<len(s); i++ {
            switch s[i] {
                case '(','{','[': {
                    stack = append(stack, s[i])
                }
                case ')','}',']': {
                    if len(stack)>0 && hash[s[i]] == stack[len(stack)-1] {
                        stack = stack[:len(stack)-1]
                    }
                }
            }
        }
        if len(stack) == 0 {
            res = true
        }
    }
    return res
}
