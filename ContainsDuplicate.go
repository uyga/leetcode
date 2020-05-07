package leetcode

func containsDuplicate(nums []int) bool {
    hashmap := make(map[int]int)
    res := false
    for i:=0; i<len(nums); i++ {
        _, key := hashmap[nums[i]]
        if (key) {
            res = true
            break
        }
        hashmap[nums[i]] = i
    }
    return res
}
