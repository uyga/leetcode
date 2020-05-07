package leetcode

func rotateArray2(nums []int, k int)  {
    if len(nums) > 1 {
        var res []int
        if k > len(nums) {
            k = k % len(nums)
        }
        res = append(res,nums[len(nums)-k:]...)
        res = append(res,nums[:len(nums)-k]...)
        copy(nums,res)
    }
}

func rotateArray(nums []int, k int)  {
    if len(nums) > 1 {
        if k > len(nums) {
            k = k % len(nums)
        }
        j:=len(nums)-1
        for i:=0;i<len(nums)/2;i++ {
            nums[i], nums[j] = nums[j], nums[i]
            j--
        }
        j=k-1
        for i:=0;i<k/2;i++ {
            nums[i], nums[j] = nums[j], nums[i]
            j--
        }
        j=len(nums)-1
        for i:=k;i<=((len(nums)-k)/2)+k-1;i++ {
            nums[i], nums[j] = nums[j], nums[i]
            j--
        }
    }
}
