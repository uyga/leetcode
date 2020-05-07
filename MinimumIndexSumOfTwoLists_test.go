package leetcode

import "testing"
import "reflect"

func TestFindRestaurant(t *testing.T) {
    cases := []struct {
        input1 []string
        input2 []string
        expected []string
    }{
        {[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},[]string{"Shogun"}},
        {[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},[]string{"KFC", "Shogun", "Burger King"},[]string{"Shogun"}},
    }
    for _,c := range cases {
        actual:=findRestaurant(c.input1, c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
