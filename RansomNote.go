package leetcode

func canConstruct(ransomNote string, magazine string) bool {
    if len(ransomNote)>len(magazine) {
        return false
    }
    hash := make(map[rune]int)
    for _,v := range magazine {
        hash[v]++
    }
    for _,v := range ransomNote {
        if _,ok := hash[v]; !ok || hash[v] == 0 {
            return false
        } else {
            hash[v]--
        }
    }
    return true
}
