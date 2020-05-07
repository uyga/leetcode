package leetcode

import "strconv"

func fizzBuzz(n int) []string {
    var res []string
    for i:=1;i<=n;i++ {
        var tmp string
        if i%3 == 0 {
            tmp = "Fizz"
        }
        if i%5 == 0 {
            tmp += "Buzz"
        }
        if len(tmp) == 0 {
            tmp = strconv.Itoa(i)
        }
        res = append(res,tmp)
    }
    return res
}
