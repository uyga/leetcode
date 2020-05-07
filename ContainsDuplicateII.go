package leetcode

func containsNearbyDuplicate2(nums []int, k int) bool {
    for i:=0;i<len(nums);i++ {
        for j:=i+1;j<=i+k;j++ {
            if len(nums) > j {
                if nums[i] == nums[j] {
                    return true
                }
            } else {
                break
            }
        }
    }
    return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
    hash := make(map[int]int)
    for i:=0;i<len(nums);i++ {
        if _,ok := hash[nums[i]]; ok {
            if i-hash[nums[i]] <= k {
                return true
            }
        }
        hash[nums[i]] = i
    }
    return false
}
