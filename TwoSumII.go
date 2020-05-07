package leetcode

func twoSumII2(nums []int, target int) []int {
    var res []int
    for i:=0;i<len(nums)-1;i++ {
        j:=i+1
        for j < len(nums) && nums[i]+nums[j] <= target {
            if nums[i]+nums[j] == target {
                res = append(res, i+1, j+1)
                return res
            }
            j++
        }
    }
    return res
}

func twoSumII(nums []int, target int) []int {
    var res []int
    if len(nums) > 0 {
        var low, high int = 0, len(nums)-1
        for low < high {
            if nums[low] + nums[high] > target {
                high --
            } else if nums[low] + nums[high] < target {
                low ++
            } else {
                res = append(res, low+1,high+1)
                break
            }
        }
    }
    return res
}
