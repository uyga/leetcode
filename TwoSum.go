package leetcode

func twoSum(nums []int, target int) []int {
    hashmap := make(map[int]int)
    for i:=0;i<len(nums);i++ {
        compliment := target - nums[i]
        val, key := hashmap[compliment]
        if key {
            return []int{i,val}
        }
        hashmap[nums[i]] = i
    }
    panic("No pair found!");
}
