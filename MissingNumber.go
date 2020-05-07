package leetcode

func missingNumber(nums []int) int {
    var sum,eth,i int
    for i=0;i<len(nums);i++ {
        eth += i
        sum += nums[i]
    }
    eth += i
    return eth - sum
}
