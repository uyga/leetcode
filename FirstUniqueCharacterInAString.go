package leetcode

func firstUniqChar2(s string) int {
    var res int = -1
    if len(s) > 0 {
        hash := make(map[rune]int)
        for k,v := range s {
            if _,ok:=hash[v]; ok {
                hash[v] = -1
            } else {
                hash[v] = k
            }
        }
        var min int = len(s)+1
        for _,v := range hash {
            if v != -1 && v < min {
                min = v
            }
        }
        if min != len(s) + 1 {
            res = min
        }
    }
    return res
}

func firstUniqChar(s string) int {
    var res int = -1
    if len(s) > 0 {
        var hash []int = make([]int, 26)
        for _,v:=range s {
            hash[v-'a']++
        }
        for i,v := range s {
            if hash[v-'a']==1 {
                return i
            }
        }
    }
    return res
}
