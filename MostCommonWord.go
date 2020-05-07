package leetcode

func mostCommonWord(paragraph string, banned []string) string {
    var word string
    var hash map[string]int = make(map[string]int)
    for i:=0;i<len(paragraph);i++ {
        if isSeparator(paragraph[i]) && len(word) > 0 {
            hash[word]++
            word = ""
        } else {
            word += filterChar(paragraph[i])
        }
    }
    if len(word) > 0 {
        hash[word]++
        word = ""
    }
    for i:=0;i<len(banned);i++ {
        delete(hash,banned[i])
    }
    max:=0
    for k,v:=range hash {
        if max < v {
            max = v
            word = k
        }
    }
    return word
}

func isSeparator(c byte) bool {
    return !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'))
}

func filterChar(c byte) string {
    var res string = ""
    if c >= 'a' && c <= 'z' {
        res = string(c)
    } else if c >= 'A' && c <= 'Z' {
        res = string(c -'A'+'a')
    }
    return res
}
