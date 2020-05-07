package leetcode

import "testing"
import "reflect"

func TestLicenseKeyFormatting(t *testing.T) {
    cases := []struct {
        input1 string
        input2 int
        expected string
    }{
        {"5F3Z-2e-9-w",4,"5F3Z-2E9W"},
        {"2-5g-3-J",2,"2-5G-3J"},
        {"as-as-12124-sdcIQ-QJJJW33",1,"A-S-A-S-1-2-1-2-4-S-D-C-I-Q-Q-J-J-J-W-3-3"},
        {"Q",4,"Q"},
        {"--a-a-a-a--",2,"AA-AA"},
        {"----------",2,""},
        {"--------EyRfCyHxyUJzhygiazYpjuDFdHvrnDwoQKQEsccLDiwhpmjueADIzqIvExbDDFnEGovAxYeszbzuTekRuWUPXRPbVKJuDQzIzzTj",16,"EYRF-CYHXYUJZHYGIAZYP-JUDFDHVRNDWOQKQE-SCCLDIWHPMJUEADI-ZQIVEXBDDFNEGOVA-XYESZBZUTEKRUWUP-XRPBVKJUDQZIZZTJ"},
    }
    for _,c := range cases {
        actual:=licenseKeyFormatting(c.input1, c.input2)
        if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("Input %#v,%#v. Expected: %#v, actual: %#v\n", c.input1, c.input2, c.expected, actual)
        }
    }
}
