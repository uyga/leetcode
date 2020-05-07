package leetcode

//O(N)
func isAnagram(s string, t string) bool {
    var res bool = false
    if len(s) == len(t) {
        var hashmap map[byte]int = make(map[byte]int)
        for i:=0;i<len(s);i++ {
            hashmap[s[i]]++
            hashmap[t[i]]--
        }
        res = true
        for _, value := range hashmap {
            if value != 0 {
                res = false
                break
            }
        }
    }
    return res
}
