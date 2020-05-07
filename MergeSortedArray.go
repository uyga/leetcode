package leetcode

func merge(nums1 []int, m int, nums2 []int, n int)  {
    if len(nums1) < n+m {
        return
    }
    var i,j,l int
    l=1
    i=m-1
    j=n-1
    for m+n-l >= 0 {
        if i<0 && j>=0 {
            nums1[m+n-l] = nums2[j]
            j--
        } else if i>=0 && j<0 {
            nums1[m+n-l] = nums1[i]
            i--
        } else if i>=0 && j>=0 {
            if nums1[i] > nums2[j] {
                nums1[m+n-l] = nums1[i]
                i--
            } else {
                nums1[m+n-l] = nums2[j]
                j--
            }
        } else {
            break
        }
        l++
    }
}
