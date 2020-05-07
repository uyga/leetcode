package leetcode

func convertToTitle(n int) string {
    var res string
    for n > 0 {
        var reminder int = n%26
        if reminder == 0 {
            res = "Z" + res
            n = n / 26 - 1
        } else {
            res = string(reminder + 'A' - 1) + res
            n = n / 26
        }
    }
    return res
}
