package leetcode

import "sort"

func singleNumber2(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    sort.Ints(nums)
    for i:=0;i<len(nums)-1;i++ {
        if nums[i]!=nums[i+1] {
            return nums[i]
        }
        i++
        if i == len(nums)-2 {
            return nums[len(nums)-1]
        }
    }
    return 0
}

func singleNumber3(nums []int) int {
    var hash map[int]bool = make(map[int]bool)
    for i:=0;i<len(nums);i++ {
        _, ok := hash[nums[i]]
        if ok {
            delete(hash, nums[i])
        } else {
            hash[nums[i]] = true
        }
    }
    for num, _ := range hash {
        return num
    }
    return 0
}

func singleNumber(nums []int) int {
    var res int
    for i:=0;i<len(nums);i++ {
        res = res ^ nums[i]
    }
    return res
}
