package leetcode

var res []string
var phone map[byte]string

func letterCombinations(digits string) []string {
    res = nil
    if len(digits) > 0 {
        phone = map[byte]string{'2':"abc",'3':"def",'4':"ghi",'5':"jkl",'6':"mno",'7':"pqrs",'8':"tuv",'9':"wxyz"}
        backtrack("",digits)
    }
    return res
}

func backtrack(combination string, next_digits string) {
    if len(next_digits) == 0 {
        res = append(res, combination)
    } else {
        letters := phone[next_digits[0]]
        for i:=0;i<len(letters);i++ {
            backtrack(combination + string(letters[i]), next_digits[1:])
        }
    }
}
