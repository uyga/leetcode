package leetcode

func max(a, b int) int {
    if a >= b {return a} else {return b}
}

func maxSubArray(nums []int) int {
    if len(nums) == 0 {return -1 << 31}
    if len(nums) == 1 {return nums[0]}
    var sum, best int = nums[0], nums[0]
    for i:=1;i<len(nums);i++ {
        sum = max(nums[i], sum + nums[i])
        best = max(best, sum)
    }
    return best
}
