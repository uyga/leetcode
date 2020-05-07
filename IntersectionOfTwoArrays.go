package leetcode

func intersect(nums1 []int, nums2 []int) []int {
    l1,l2 := len(nums1), len(nums2)
    if l1 > l2 {
        l1,l2,nums1,nums2 = l2,l1,nums2,nums1
    }
    var res []int
    var hash map[int]int = make(map[int]int, l1)
    for i:=0;i<l1;i++ {
        hash[nums1[i]]++
    }
    for i:=0;i<l2;i++ {
        if _,ok := hash[nums2[i]]; ok {
            if hash[nums2[i]] > 0 {
                res = append(res,nums2[i])
                hash[nums2[i]]--
            }
        }
    }
    return res
}
