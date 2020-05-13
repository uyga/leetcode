package leetcode

func numJewelsInStones(J string, S string) int {
    var res int
    hash := make([]byte,128)
    for i:=0;i<len(J);i++ {
        hash[J[i]]++
    }
    for i:=0;i<len(S);i++ {
        res+=int(hash[S[i]])
    }
    return res
}
