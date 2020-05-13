package leetcode

func lemonadeChange(bills []int) bool {
    fives,tens,twenties := 0,0,0
    for i:=0;i<len(bills);i++ {
        switch bills[i] {
            case 20: {
                twenties++
                if tens > 0 {
                    tens--
                } else {
                    fives--
                    fives--
                }
                fives--
            }
            case 10: {
                tens++
                fives--
            }
            case 5: {
                fives++
            }
        }
        for _,v:=range []int{fives,tens,twenties} {
            if v < 0 {
                return false
            }
        }
    }
    return true
}
