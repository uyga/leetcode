package leetcode

func licenseKeyFormatting(S string, K int) string {
    var res string
    var block []byte = make([]byte, K)
    var j int = K
    for i:=len(S)-1;i>=0;i-- {
        if S[i] != '-' {
            j--
            block[j] = toUpper(S[i])
            if j == 0 || i == 0 {
                res = "-" + string(block[j:]) + res
                j = K
                block = make([]byte,K)
            }
        }
    }
    if j != 0 && j != K {
        res = string(block[j:]) + res
    }
    if len(res)>0 && res[0] == '-' {
        res = res[1:]
    }
    return res
}

func toUpper(c byte) byte {
    if c >= 'a' && c <= 'z' {
        c = c - 'a' + 'A'
    }
    return c
}
