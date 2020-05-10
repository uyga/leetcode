package leetcode

import "testing"

func TestMyHashMap(t *testing.T) {
    obj := MyHashMapConstructor();
    if obj.Get(1) != -1 {
        t.Errorf("Failed on Get when empty map.")
    }
    obj.Put(1,1)
    if obj.Get(1) != 1 {
        t.Errorf("Failed on Get after Put.")
    }
    for i:=0;i<10000;i++ {
        obj.Put(i,i+5)
    }
    if obj.Get(6962) == -1 {
        t.Errorf("Failed on Get after Put.")
    }
    if obj.Get(0) == -1 {
        t.Errorf("Failed on Get after Put.")
    }
    if obj.Get(1000000) != -1 {
        t.Errorf("Failed on Get when no value in the set")
    }
    obj.Remove(5999)
    obj.Remove(1999)
    obj.Remove(9999)
    obj.Remove(1000000)
    obj.Remove(0)
    if obj.Get(0) != -1 {
        t.Errorf("Failed on Get after Remove.")
    }
}
