package leetcode

func reverseVowels2(s string) string {
    res:= []rune(s)
    i := 0
    j := len(s)-1
    for i <= j {
        for i<=j && !isVowel(res[i]) {i++}
        for j>=i && !isVowel(res[j]) {j--}
        if i<len(s) && isVowel(res[i]) && isVowel(res[j]) {
            res[i],res[j] = res[j],res[i]
            i++
            j--
        }
    }
    return string(res)
}

func reverseVowels(s string) string {
    var vowels []int
    for k,v := range s {
        if isVowel(v) {
            vowels = append(vowels, k)
        }
    }
    res := []byte(s)
    i:=0
    j:=len(vowels)-1
    for i<=j {
        res[vowels[i]],res[vowels[j]] = res[vowels[j]],res[vowels[i]]
        i++
        j--
    }
    return string(res)
}

func isVowel(c rune) bool {
    switch c {
        case 'a':
            return true
        case 'e':
            return true
        case 'i':
            return true
        case 'o':
            return true
        case 'u':
            return true
        case 'A':
            return true
        case 'E':
            return true
        case 'I':
            return true
        case 'O':
            return true
        case 'U':
            return true
    }
    return false
}
