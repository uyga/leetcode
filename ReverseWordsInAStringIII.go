package leetcode

func reverseWords(s string) string {
    var data []byte = []byte(s)
    i:=0
    for i<len(data) {
        for i<len(data) && data[i] == ' ' {
            i++
        }
        pos:=i
        for pos<len(data) && data[pos] != ' ' {
            pos++
        }
        if i != pos-1 {
            begin:=i
            end:=pos-1
            for begin <= end {
                data[begin], data[end] = data[end], data[begin]
                begin++
                end--
            }
            i = pos-1
        }
        i++
    }
    return string(data)
}
