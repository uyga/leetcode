package leetcode

import "testing"
import "reflect"

func TestIntToRoman(t *testing.T) {
    cases := []struct {
        input int
        expected string
    }{
        {1,"I"},
        {3,"III"},
        {4,"IV"},
        {8,"VIII"},
        {9,"IX"},
        {10,"X"},
        {11,"XI"},
        {14,"XIV"},
        {58,"LVIII"},
        {1994,"MCMXCIV"},
        {1000,"M"},
        {1500,"MD"},
    }
    for _,c := range cases {
        actual:=intToRoman(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
