package leetcode

func toLowerCase(str string) string {
    res := []byte(str)
    for i:=0;i<len(res);i++ {
        if res[i] >= 'A' && res[i] <= 'Z' {
            res[i] = res[i]-'A'+'a'
        }
    }
    return string(res)
}
