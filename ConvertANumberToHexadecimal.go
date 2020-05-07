package leetcode

func toHex(num int) string {
    hash := map[uint32]rune{0:'0',1:'1',2:'2',3:'3',4:'4',5:'5',6:'6',7:'7',8:'8',9:'9',10:'a',11:'b',12:'c',13:'d',14:'e',15:'f'}
    var res string
    if num == 0 {
        res = "0"
    } else {
        var n uint32 = uint32(num)
        for n > 0 {
            rem := n % 16
            res = string(hash[rem]) + res
            n /= 16
        }
    }
    return res
}
