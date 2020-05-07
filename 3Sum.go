package leetcode

import "sort"

func threeSum(nums []int) [][]int {
    var res [][]int
    if len(nums) > 2 {
        sort.Ints(nums)
        for i:=0;i<len(nums);i++ {
            var low, high, sum int = i+1, len(nums)-1, 0
                for low < high {
                    sum = nums[i] + nums[low] + nums[high]
                    if sum > 0 {
                        for low<high && nums[high] == nums[high-1] {high--}
                        high--
                    } else if sum < 0 {
                        for low<high && nums[low] == nums[low+1] {low++}
                        low++
                    } else {
                        res = append(res, []int{nums[i], nums[low], nums[high]})
                        for low<high && nums[high] == nums[high-1] {high--}
                        for low<high && nums[low] == nums[low+1]{low++}
                        high--
                        low++
                        for nums[i] == nums[i+1] && i < len(nums)-2 {i++}
                    }
                }
        }
    }
    return res
}

