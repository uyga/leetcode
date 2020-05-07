package leetcode

func removeDuplicates2(nums []int) int {
    var res []int
    var n int
    if len(nums) > 0 {
        res = append(res, nums[0])
        n++
        for i:=0;i<len(nums)-1;i++ {
            if nums[i] != nums[i+1] {
                res = append(res, nums[i+1])
                n++
            }
        }
    }
    copy(nums,res)
    return n
}

func removeDuplicates(nums []int) int {
    var n int
    if len(nums) > 0 {
        n++
        for i:=0;i<len(nums)-1;i++ {
            if nums[i] != nums[i+1] {
                nums[n] = nums[i+1]
                n++
            }
        }
    }
    return n
}
