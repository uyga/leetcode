package leetcode

import "sort"

func thirdMax(nums []int) int {
    sort.Ints(nums)
    var res int
    if len(nums) > 0 {
        if len(nums) < 3 {
            res = nums[len(nums)-1]
        } else {
            var cnt,i int
            for i=len(nums)-1;i>0;i-- {
                if nums[i] != nums[i-1] {
                    cnt++
                }
                if cnt == 3 {
                    break
                }
            }
            if cnt < 2 {
                res = nums[len(nums)-1]
            } else {
                res = nums[i]
            }
        }
    }
    return res
}
