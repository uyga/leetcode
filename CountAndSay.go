package leetcode

var countAndSayCache map[int]string

func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    if countAndSayCache == nil {
        countAndSayCache = make(map[int]string)
    }
    var res string
    if _,ok := countAndSayCache[n-1]; !ok {
        countAndSayCache[n-1] = countAndSay(n-1)
    }
    i:=0
    for i<len(countAndSayCache[n-1]) {
        h:=countAndSayCache[n-1][i]
        cnt := 0
        for i<len(countAndSayCache[n-1]) && countAndSayCache[n-1][i]==h {
            cnt++
            i++
        }
        res += string(cnt+'0') + string(h)
    }
    return res
}
