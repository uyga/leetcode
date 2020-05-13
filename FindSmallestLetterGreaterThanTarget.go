package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
    var res byte = 255
    for i:=0;i<len(letters);i++ {
        if letters[i] > target {
            if letters[i] < res {
                res = letters[i]
            }
        }
    }
    if res == 255 {
        for i:=0;i<len(letters);i++ {
            if letters[i] < target {
                if letters[i] < res {
                    res = letters[i]
                }
            }
        }
    }
    return res
}
