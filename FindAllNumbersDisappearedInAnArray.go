package leetcode

import "math"

func findDisappearedNumbers(nums []int) []int {
    var res []int
    for i:=0;i<len(nums);i++ {
        el := int(math.Abs(float64(nums[i]))) - 1
        if nums[el] > 0 {
            nums[el] *= -1
        }
    }
    for i:=0;i<len(nums);i++ {
        if nums[i] > 0 {
            res = append(res, i+1)
        }
    }
    return res
}
