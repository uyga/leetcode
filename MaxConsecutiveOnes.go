package leetcode

func findMaxConsecutiveOnes(nums []int) int {
    var i,j int = 0,len(nums)-1
    var b,e int
    var max int
    for i<=j {
        for i<=j && nums[i] == 0 {
            i++
        }
        b = i
        for i<=j && nums[i] == 1 {
            i++
        }
        e = i
        if max < e-b {
            max = e-b
        }
    }
    return max
}
