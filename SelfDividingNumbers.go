package leetcode

import "strconv"

func selfDividingNumbers(left int, right int) []int {
    var res []int
    for i:=left;i<=right;i++ {
        str:=strconv.Itoa(i)
        cnt:=0
        for j:=len(str)-1;j>=0;j-- {
            if str[j]=='0' {
                break
            }
            if i % int(str[j]-'0') == 0 {
                cnt++
            }
        }
        if cnt == len(str) {
            res = append(res,i)
        }
    }
    return res
}
