package leetcode

func add(a byte, b byte, c byte) (res byte, remainder byte) {
    if a == '1' && b == '1' {
        res = '0' | c
        remainder = '1'
    } else if a == '0' && b == '0' {
        res = '0' | c
        remainder = '0'
    } else {
        if c == '0' {
            res = '1'
            remainder = '0'
        } else {
            res = '0'
            remainder = '1'
        }
    }
    return res, remainder
}

func addBinary(a string, b string) string {
    var remainder byte = '0'
    var curr, opp *string = &a, &b
    if len(*curr) < len(*opp) {
        curr, opp = opp, curr
    }
    var res []byte = make([]byte, len(*curr)+1)
    for i:=len(*curr)-1;i>=0;i-- {
        j:=i - (len(*curr) - len(*opp));
        if (j >= 0) {
            res[i+1], remainder = add(string(*curr)[i], string(*opp)[j], remainder)
        } else {
            res[i+1], remainder = add(string(*curr)[i], remainder, '0')
        }
    }
    if remainder == '1' {
        res[0] = '1'
        return string(res)
    }
    return string(res[1:])
}
