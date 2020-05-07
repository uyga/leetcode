package leetcode

func isBadVersion(version int) bool {
    return version >= 4
}

func firstBadVersion(n int) int {
    var left, right int = 1, n
    for left < right {
        var mid int = (right - left)/2 + left
        if isBadVersion(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}
