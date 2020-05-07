package leetcode

func moveZeroes(nums []int)  {
    var j int
    for i:=0;i<len(nums);i++ {
        if nums[i] == 0 && i<len(nums)-1 {
            j = i+1
            for j<len(nums)-1 && nums[j] == 0 {
                j++
            }
            nums[i],nums[j] = nums[j],nums[i]
            if j >= len(nums)-1 {
                break
            }
        }
    }
}
