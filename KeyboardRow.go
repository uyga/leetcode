package leetcode

func findWords(words []string) []string {
    var res []string
    h1 := map[byte]bool{'q':true,'w':true,'e':true,'r':true,'t':true,'y':true,'u':true,'i':true,'o':true,'p':true}
    h2 := map[byte]bool{'a':true,'s':true,'d':true,'f':true,'g':true,'h':true,'j':true,'k':true,'l':true}
    h3 := map[byte]bool{'z':true,'x':true,'c':true,'v':true,'b':true,'n':true,'m':true}
    var dic *map[byte]bool
    for _,word := range words {
        if len(word) > 0 {
            if _,ok := h1[toLower(word[0])]; ok {
                dic = &h1
            } else if _,ok := h2[toLower(word[0])]; ok {
                dic = &h2
            } else {
                dic = &h3
            }
        }
        var i int
        for i=0;i<len(word);i++ {
            if _,ok := map[byte]bool(*dic)[toLower(word[i])]; !ok {
                break
            }
        }
        if i == len(word) {
            res = append(res, word)
        }
    }
    return res
}

func toLower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        c = c - 'A' + 'a'
    }
    return c
}
