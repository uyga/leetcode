package leetcode

func rob(nums []int) int {
    even, odd := 0, 0
    for i:=0;i<len(nums);i++ {
        if i % 2 == 0 {
            even += nums[i]
            if even < odd {
                even = odd
            }
        } else {
            odd += nums[i]
            if odd < even {
                odd = even
            }
        }
    }
    if odd > even {
        even = odd
    }
    return even
}
