package leetcode

func compress(chars []byte) int {
    var curr, size, pos int
    for curr < len(chars) {
        size = 1
        for curr < len(chars)-1 && chars[curr] == chars[curr+1] {
            size++
            curr++
        }
        chars[pos] = chars[curr]
        if size > 1 {
            pos++
            if size >= 1000 {
                chars[pos+3] = byte(size%10)+'0'
                size /= 10
                chars[pos+2] = byte(size%10)+'0'
                size /= 10
                chars[pos+1] = byte(size%10)+'0'
                size /= 10
                chars[pos] = byte(size%10)+'0'
                pos += 3
            } else if size >= 100 {
                chars[pos+2] = byte(size%10)+'0'
                size /= 10
                chars[pos+1] = byte(size%10)+'0'
                size /= 10
                chars[pos] = byte(size)+'0'
                pos+=2
            } else if size >= 10 {
                chars[pos+1] = byte(size%10)+'0'
                size /= 10
                chars[pos] = byte(size)+'0'
                pos++
            } else {
                chars[pos] = byte(size)+'0'
            }
        }
        pos++
        curr++
    }
    return pos
}
