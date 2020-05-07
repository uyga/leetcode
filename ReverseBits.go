package leetcode

import "math"

func reverseBits2(num uint32) uint32 {
    var res uint32
    var i,j float64
    j=31
    for i=0;i<32;i++ {
        if num & uint32(math.Pow(2,i)) != 0 {
            res |= uint32(math.Pow(2,j))
        }
        j--
    }
    return res
}

func reverseBits3(num uint32) uint32 {
    var res,remainder uint32
    for i:=0;i<32;i++ {
        remainder = num % 2
        num /= 2
        res = res * 2 + remainder
    }
    return res
}

func reverseBits(num uint32) uint32 {
    var res uint32
    for i:=31;i>=0;i-- {
        if num & (1<<i) != 0 {
            res = res | (1<<(31-i))
        }
    }
    return res
}
