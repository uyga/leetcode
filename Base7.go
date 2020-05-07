package leetcode

func convertToBase7(num int) string {
    if num == 0 {
        return "0"
    }
    var res string
    var sign int = 1
    if num < 0 {
        num = num * -1
        sign = -1
    }
    for num > 0 {
        res = string(byte(num%7)+'0') + res
        num /= 7
    }
    if sign < 0 {
        res = "-" + res
    }
    return res
}
