package leetcode

func findTheDifference2(s string, t string) byte {
    var hash []int = make([]int,26)
    for _,v := range s {
        hash[v-'a']++
    }
    for _,v:= range t {
        if hash[v-'a'] == 0 {
            return byte(v)
        } else {
            hash[v-'a']--
        }
    }
    return 0
}

func findTheDifference(s string, t string) byte {
    var e byte
    for k,_ := range s {
        e = s[k] ^ t[k]
    }
    e = e ^ t[len(t)-1]
    return e
}
