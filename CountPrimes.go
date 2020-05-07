package leetcode

func countPrimes(n int) int {
    var primes []int
    for i:=2;i<n;i+=1 {
        v := 1
        for _,prime := range primes {
            if i == prime || i%prime == 0 {
                v = 0
                break
            }
            if prime*prime > i {
                break
            }
        }
        if v==1 {
            primes = append(primes, i)
        }
    }
    return len(primes)
}
