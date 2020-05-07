package leetcode

import "testing"
import "reflect"

func TestMerge(t *testing.T) {
    nums1:=[]int{1,2,3,0,0,0,0}
    nums2:=[]int{2,5,6,8}
    m:=3
    n:=4
    merge(nums1,m,nums2,n)
    if !reflect.DeepEqual(nums1, []int{1,2,2,3,5,6,8}) {
        t.Errorf("The result is not expected: %#v\n", nums1)
    }
}
