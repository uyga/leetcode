package leetcode

func addStrings2(num1 string, num2 string) string {
    if len(num1) == 0 && len(num2) == 0 {
        return "0"
    }
    var p1,p2 *string = &num1, &num2
    var res string
    var tmp byte
    if len(*p2) > len(*p1) {
        p1,p2 = p2,p1
    }
    var rem byte
    var i,j int
    for i=len(*p1)-1;i>=0;i-- {
        tmp = rem + []byte(*p1)[i]-'0'
        j = i - (len(*p1) - len(*p2))
        if j >= 0 {
            tmp += []byte(*p2)[j]-'0'
        }
        rem = 0
        if tmp > 9 {
            tmp = tmp - 10
            rem = 1
        }
        res = string(tmp+'0') + res
        if i == 0 && rem != 0 {
            res = string(rem+'0') + res
        }
    }
    return string(res)
}

func addStrings(num1 string, num2 string) string {
    if len(num1) == 0 && len(num2) == 0 {
        return "0"
    }
    l1 := len(num1)-1
    l2 := len(num2)-1
    if l1 < l2 {
        l1,l2,num1,num2 = l2,l1,num2,num1
    }
    var res []byte = make([]byte,l1+2)
    var rem  byte
    for i:=l1;i>=0;i-- {
        tmp := num1[i]-'0' + rem
        j:=i-(l1-l2)
        if j >=0 {
            tmp += num2[j]-'0'
        }
        rem = 0
        if tmp > 9 {
            tmp = tmp - 10
            rem = 1
        }
        res[i+1]=tmp+'0'
    }
    if rem != 0 {
        res[0]=rem+'0'
        return string(res)
    }
    return string(res[1:])
}
