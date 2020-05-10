package leetcode

import "testing"

func TestMyHashSet(t *testing.T) {
    obj := MyHashSetConstructor();
    if obj.Contains(1) {
        t.Errorf("Failed on Contains when empty set.")
    }
    obj.Add(1)
    if !obj.Contains(1) {
        t.Errorf("Failed on Contains after addition.")
    }
    for i:=0;i<10000;i++ {
        obj.Add(i)
    }
    if !obj.Contains(6962) {
        t.Errorf("Failed on Contains after addition.")
    }
    if !obj.Contains(0) {
        t.Errorf("Failed on Contains after addition.")
    }
    if obj.Contains(1000000) {
        t.Errorf("Failed on Contains when no value in the set")
    }
    obj.Remove(5999)
    obj.Remove(1999)
    obj.Remove(9999)
    obj.Remove(1000000)
    obj.Remove(0)
    if obj.Contains(0) {
        t.Errorf("Failed on Contains after remove.")
    }
}
