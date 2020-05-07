package leetcode

func removeElement(nums []int, val int) int {
    var n int
    if len(nums) > 0 {
        for i:=0;i<len(nums);i++ {
            if nums[i] != val {
                nums[n] = nums[i]
                n++
            }
        }
    }
    return n
}
