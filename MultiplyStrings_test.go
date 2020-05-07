package leetcode

import "testing"
import "reflect"

func TestMultiply(t *testing.T) {
    cases := []struct {
        input []string
        expected string
    }{
        {[]string{"2","3"},"6"},
        {[]string{"123","456"},"56088"},
        {[]string{"99","99"},"9801"},
        {[]string{"1234","5678910"},"7007774940"},
        {[]string{"65536","65536"},"4294967296"},
        {[]string{"123456789","987654321"},"121932631112635269"},
        {[]string{"0","3"},"0"},
        {[]string{"2","0"},"0"},
        {[]string{"0","0"},"0"},
        {[]string{"2147483647","2147483647"},"4611686014132420609"},
        {[]string{"461168601413242060946116860141324206094611686014132420609","461168601413242060946116860141324206094611686014132420609"},"212676478929445727412523898482768454290378014907259233583525011038765270719061244853286703190356998860269687930881"},
    }
    for _,c := range cases {
        actual:=multiply(c.input[0],c.input[1])
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v. Expected: %#v, actual: %#v\n", c.input, c.expected, actual)
        }
    }
}
