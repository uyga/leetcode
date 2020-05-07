package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
        return true
    }
    i:=0;
    l:=len(flowerbed)
    for i<l {
        for i<l && flowerbed[i] == 1 {
            i++
        }
        for i<l && flowerbed[i] == 0 {
            if (i>1 && flowerbed[i-1] == 0 && i<l-2 && flowerbed[i+1] == 0) ||
            (i == 0 && i<l-1 && flowerbed[i+1] == 0) ||
            (i == l-1 && i>0 && flowerbed[i-1] == 0) ||
            (l == 1) {
                n--
                if n == 0 {
                    return true
                }
                i++
            }
            i++
        }
    }
    return n<=0
}
