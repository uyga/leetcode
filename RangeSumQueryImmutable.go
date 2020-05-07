package leetcode

type NumArray struct {
    sums []int
}


func NumArrayConstructor(nums []int) NumArray {
    var na NumArray
    sum := 0
    for i:=0;i<len(nums);i++{
        sum += nums[i]
        na.sums=append(na.sums,sum)
    }
    return na
}


func (this *NumArray) SumRange(i int, j int) int {
    if i == 0 {
        return this.sums[j]
    } else {
        return this.sums[j] - this.sums[i-1]
    }
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
