package leetcode

func searchInsert(nums []int, target int) int {
    return searchInsertBisect(nums, target, 0, len(nums)-1)
}

func searchInsertBisect(nums []int, target int, start int, end int) int {
    var res int
    if end - start >= 0 {
        if nums[start] >= target {
            res = start
        } else if nums[end] == target {
            res = end
        } else if nums[end] < target {
            res = end + 1
        } else {
            var mid int = start + ((end - start) / 2)
            if nums[mid] > target {
                res = searchInsertBisect(nums, target, start, mid)
            } else if nums[mid] < target {
                res = searchInsertBisect(nums, target, mid+1, end)
            } else {
                res = mid
            }
        }
    }
    return res
}
