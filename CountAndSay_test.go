package leetcode

import "testing"
import "reflect"

func TestCountAndSay(t *testing.T) {
    cases := []struct {
        input int
        expected string
    }{
        {1,"1"},
        {2,"11"},
        {3,"21"},
        {4,"1211"},
        {5,"111221"},
        {9,"31131211131221"},
        {10,"13211311123113112211"},
        {20,"11131221131211132221232112111312111213111213211231132132211211131221131211221321123113213221123113112221131112311332211211131221131211132211121312211231131112311211232221121321132132211331121321231231121113112221121321133112132112312321123113112221121113122113121113123112112322111213211322211312113211"},
    }
    for _,c := range cases {
        actual:=countAndSay(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
