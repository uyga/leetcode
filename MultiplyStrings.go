package leetcode

func multiply(num1 string, num2 string) string {
    if num1=="0" || num2 == "0" {
        return "0"
    }
    var res []byte = make([]byte, len(num1)+len(num2))
    for i:=len(num1)-1;i>=0;i-- {
        for j:=len(num2)-1;j>=0;j-- {
            res[i+j+1] += (num1[i]-'0') * (num2[j]-'0')
            for k:=i+j+1;res[k]>9;k-- {
                res[k-1] += res[k] / 10
                res[k] = res[k] % 10
            }
        }
    }
    var s string
    if res[0] != 0 {
        s = string(res[0]+'0')
    }
    for i:=1;i<len(res);i++ {
        s += string(res[i]+'0')
    }
    return s
}

func multiply2(num1 string, num2 string) string {
    if (len(num1) == 1 && num1[0] == '0') || (len(num2) == 1 && num2[0] == '0') {
        return "0"
    }
    l1,l2 := len(num1)-1,len(num2)-1
    if l2 > l1 {
        l1,l2,num1,num2 = l2,l1,num2,num1
    }
    tmp := 0
    rem := 0
    var total [][]int
    for i:=l1;i>=0;i-- {
        var buff []int
        for j:=0;j<l1-i;j++ {
            buff = append(buff,0)
        }
        for j:=l2;j>=0;j-- {
            tmp = int(num1[i]-'0') * int(num2[j]-'0') + rem
            if tmp > 9 {
                rem = tmp/10
                tmp = tmp%10
            } else {
                rem = 0
            }
            buff = append(buff,tmp)
        }
        if rem != 0 {
            buff = append(buff, rem)
            rem = 0
        }
        total = append(total,buff)
    }
    return sumArrays(total)
}

func sumArrays(a [][]int) string {
    var out []byte
    var res []int = make([]int, len(a[len(a)-1])+1)
    var rem,i,j int
    for i=len(a)-1;i>=0;i-- {
        for j=0;j<len(a[i]);j++ {
            res[j] = res[j]+a[i][j] + rem
            if res[j] > 9 {
                rem = res[j]/10
                res[j] = res[j]%10
            } else {
                rem = 0
            }
        }
        if rem != 0 {
            res[j] += rem
            for res[j] > 9 {
                rem = res[j]/10
                res[j] = res[j]%10
                j++
                res[j] += rem
            }
            rem = 0
        }
    }
    l := len(res)-1
    if res[l] == 0 {
        l--
    }
    for i=l;i>=0;i-- {
        out = append(out, byte(res[i]+'0'))
    }
    return string(out)
}
