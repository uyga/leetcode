package leetcode

//O(N)
func majorityElement2(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    var res int
    var hash map[int]int = make(map[int]int)
    i,j:=0,len(nums)-1
    for i <= j {
        hash[nums[i]]++
        hash[nums[j]]++
        if hash[nums[i]] > len(nums)/2 {
            return nums[i]
        }
        if hash[nums[j]] > len(nums)/2 {
            return nums[j]
        }
        i++
        j--
    }
    return res
}

//Boyer-Moore voting algorithm
//O(N)
func majorityElement(nums []int) int {
    var res,cnt int
    if len(nums) == 0 {return 0}
    for _,num := range nums {
        if cnt == 0 {
            res, cnt = num, 1
        } else {
            if res == num {
                cnt++
            } else {
                cnt--
            }
        }
    }
    return res
}
