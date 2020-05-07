package leetcode

func longestCommonPrefix(strs []string) string {
    if len(strs) > 0 {
        var res int
        if len(strs[0]) > 0 {
            var cnt,i,j int
            for j=0;;j++ {
                if j >= len(strs[0]) {
                    break
                }
                for i=0;i<len(strs);i++ {
                    if len(strs[i]) > 0 && j < len(strs[i]) {
                        if strs[0][j] == strs[i][j] {
                            cnt++
                        }
                    } else {
                        break
                    }
                }
                if cnt == len(strs) {
                    res ++
                    cnt = 0
                } else {
                    break
                }
            }
        }
        return strs[0][:res]
    }
    return ""
}
