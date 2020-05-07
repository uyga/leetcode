package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    var res []int
    hash := make(map[int]int, len(nums2))
    for i:=0;i<len(nums2);i++ {
        hash[nums2[i]]=i
    }
    for i:=0;i<len(nums1);i++ {
        if _,ok := hash[nums1[i]]; !ok {
            res = append(res, -1)
        } else {
            r := -1
            j := hash[nums1[i]]+1
            for j<len(nums2) {
                if nums2[j] > nums1[i] {
                    r = nums2[j]
                    break
                }
                j++
            }
            res = append(res, r)
        }
    }
    return res
}
