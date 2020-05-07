package leetcode

func romanToInt(s string) int {
    var res,curr,next int
    if len(s) > 0 {
        numbers := map[byte]int{'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
        if len(s) == 1 {
            res = numbers[s[0]]
        } else {
            for i:=0;i<len(s)-1;i++ {
                curr = numbers[s[i]]
                next = numbers[s[i+1]]
                if curr >= next {
                    res += curr
                } else {
                    res -= curr
                }
            }
            res += next
        }
    }
    return res
}
