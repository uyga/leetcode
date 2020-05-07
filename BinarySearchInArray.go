package leetcode

func binarySearchInArray(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    return bisectInArray(0,len(nums)-1,nums,target)
}

func bisectInArray(start,end int, nums []int, target int) int {
    if nums[start] == target {
        return start
    }
    if nums[end] == target {
        return end
    }
    if start >= end {
        return -1
    }
    mid:=start+(end-start)/2
    if nums[mid] > target {
        return bisectInArray(start,mid,nums,target)
    } else if nums[mid] < target {
        return bisectInArray(mid+1,end,nums,target)
    } else {
        return mid
    }
}
