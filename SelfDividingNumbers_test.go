package leetcode

import "testing"
import "reflect"

func TestSelfDividingNumbers(t *testing.T) {
    cases := []struct {
        input []int
        expected []int
    }{
        {[]int{1,22},[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
        {[]int{1,1000},[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22, 24, 33, 36, 44, 48, 55, 66, 77, 88, 99, 111, 112, 115, 122, 124, 126, 128, 132, 135, 144, 155, 162, 168, 175, 184, 212, 216, 222, 224, 244, 248, 264, 288, 312, 315, 324, 333, 336, 366, 384, 396, 412, 424, 432, 444, 448, 488, 515, 555, 612, 624, 636, 648, 666, 672, 728, 735, 777, 784, 816, 824, 848, 864, 888, 936, 999}},
        {[]int{50,100},[]int{55, 66, 77, 88, 99}},
        {[]int{0,1},[]int{1}},
    }
    for _,c := range cases {
        actual:=selfDividingNumbers(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
