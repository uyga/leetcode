package leetcode

func findRestaurant(list1 []string, list2 []string) []string {
    l1, l2 := len(list1), len(list2)
    if l1 < l2 {
        l1,l2,list1,list2 = l2,l1,list2,list1
    }
    var hash map[string]int = make(map[string]int,l1)
    var minsum int = l1*l1
    for i:=0;i<l1;i++ {
        hash[list1[i]] = i
    }
    for i:=0;i<l2;i++ {
        if _,ok := hash[list2[i]]; ok {
            sum := hash[list2[i]] + i + l1
            hash[list2[i]] = sum
            if sum < minsum {
                minsum = sum
            }
        }
    }
    var res []string
    for k,v := range hash {
        if v == minsum {
            res = append(res, k)
        }
    }
    return res
}
