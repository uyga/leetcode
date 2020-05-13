package leetcode

import "strconv"

func maximum69Number(num int) int {
    n := []byte(strconv.Itoa(num))
    for i:=0;i<len(n);i++ {
        if n[i] == '6' {
            n[i] = '9'
            break
        }
    }
    res,_ := strconv.Atoi(string(n))
    return res
}
