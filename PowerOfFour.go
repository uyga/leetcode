package leetcode

import "math/bits"

func isPowerOfFour2(num int) bool {
    if num <= 0 {
        return false
    }
    return num & (num - 1) == 0 && bits.TrailingZeros32(uint32(num)) % 2 == 0
}

func isPowerOfFour(num int) bool {
    if num <= 0 {
        return false
    }
    return num & (num - 1) == 0 && num & 0x55555555 != 0
}
