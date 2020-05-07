package leetcode

func isPowerOfTwo2(n int) bool {
    if n <= 0 {
        return false
    }
    if n==1 || n%2 == 0 {
        for i:=1;i<=n;i*=2 {
            if i<n && (i&n) != 0 {
                return false
            }
        }
        return true
    }
    return false
}

func isPowerOfTwo3(n int) bool {
    if n<=0 {
        return false
    }
    for n>1 {
        if n%2 != 0 {
            return false
        }
        n/=2
    }
    return true
}

func isPowerOfTwo(n int) bool {
    if n <=0 {
        return false
    }
    return n & (n-1) == 0
}
