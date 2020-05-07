package leetcode

func productExceptSelf(nums []int) []int {
    answer := make([]int, len(nums))
    answer[0] = 1
    for i:=1;i<len(nums);i++ {
        answer[i] = nums[i-1]*answer[i-1]
    }
    var r int = 1
    for i:=len(nums)-1;i>=0;i-- {
        answer[i] = answer[i] * r
        r = nums[i] * r
    }
    return answer
}
