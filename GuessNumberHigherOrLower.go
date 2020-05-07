package leetcode

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(n int) int {
    x := 3
    if n < x {
        return -1
    } else if n > x {
        return 1
    }
    return 0
}

func guessNumber(n int) int {
    var low,high,mid int = 1,n,0
    for low<high {
        mid = (high-low)/2+low
        switch guess(mid) {
            case -1:
                high = mid
            case 1:
                low = mid+1
            default:
                return mid
        }
    }
    return low
}
