package leetcode

func intToRoman(num int) string {
    var res string
    dict := []struct{
        num int
        val string
    }{
        {1,"I"},
        {4,"IV"},
        {5,"V"},
        {9,"IX"},
        {10,"X"},
        {40,"XL"},
        {50,"L"},
        {90,"XC"},
        {100,"C"},
        {400,"CD"},
        {500,"D"},
        {900,"CM"},
        {1000,"M"},
    }
    for i:=len(dict)-1;i>=0;i-- {
        for num >= dict[i].num {
            res += dict[i].val
            num -= dict[i].num
            if num == 0 {
                return res
            }
        }
    }
    return res
}
