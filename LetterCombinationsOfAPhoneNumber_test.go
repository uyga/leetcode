package leetcode

import "testing"
import "reflect"

func TestLetterCombinations(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    }{
        {"23",[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
        {"2",[]string{"a", "b", "c"}},
        {"22",[]string{"aa","ab","ac","ba","bb","bc","ca","cb","cc"}},
        {"6545",[]string{"mjgj","mjgk","mjgl","mjhj","mjhk","mjhl","mjij","mjik","mjil","mkgj","mkgk","mkgl","mkhj","mkhk","mkhl","mkij","mkik","mkil","mlgj","mlgk","mlgl","mlhj","mlhk","mlhl","mlij","mlik","mlil","njgj","njgk","njgl","njhj","njhk","njhl","njij","njik","njil","nkgj","nkgk","nkgl","nkhj","nkhk","nkhl","nkij","nkik","nkil","nlgj","nlgk","nlgl","nlhj","nlhk","nlhl","nlij","nlik","nlil","ojgj","ojgk","ojgl","ojhj","ojhk","ojhl","ojij","ojik","ojil","okgj","okgk","okgl","okhj","okhk","okhl","okij","okik","okil","olgj","olgk","olgl","olhj","olhk","olhl","olij","olik","olil"}},
    }
    for _,c := range cases {
        actual:=letterCombinations(c.input)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
