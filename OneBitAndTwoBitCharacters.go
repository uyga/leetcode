package leetcode

func isOneBitCharacter(bits []int) bool {
    if len(bits) == 1 {
        return true
    }
    l:=len(bits)-1
    i:=0
    for i<l {
        for i<l && bits[i] == 0 {
            i++
        }
        for i<l && bits[i] == 1 {
            i+=2
        }
    }
    return i==l
}
