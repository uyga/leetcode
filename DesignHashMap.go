package leetcode

type Unit struct {
    Key int
    Val int
    Next *Unit
}

type HashMapTens struct {
    bucket0 *HashMapTens
    bucket1 *HashMapTens
    bucket2 *HashMapTens
    bucket3 *HashMapTens
    bucket4 *HashMapTens
    bucket5 *HashMapTens
    bucket6 *HashMapTens
    bucket7 *HashMapTens
    bucket8 *HashMapTens
    bucket9 *HashMapTens
    Next *HashMapTens
    Data *Unit
}

type MyHashMap struct {
    HashMapTens
}


/** Initialize your data structure here. */
func MyHashMapConstructor() MyHashMap {
    return MyHashMap{}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    t := this.GetBucket(key, true)
    if t.Data == nil {
        t.Data = new(Unit)
        t.Data.Key = key
        t.Data.Val = value
    } else {
        d := t.Data
        for d.Next != nil && d.Key != key {
            d = d.Next
        }
        if d.Next == nil && d.Key != key {
            d.Next = new(Unit)
            d.Next.Key = key
            d.Next.Val = value
        }
        if d.Key == key {
            d.Val = value
        }
    }
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    var res int = -1
    t := this.GetBucket(key, false)
    d := t.Data
    for d != nil {
        if d.Key == key {
            res = d.Val
            break
        }
        d = d.Next
    }
    return res
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    t := this.GetBucket(key, false)
    d := t.Data
    prev := d
    for d != nil {
        if d.Key == key {
            if prev == d {
                t.Data = d.Next
            }
            if d.Next != nil {
                prev.Next = d.Next
            } else {
                prev.Next = nil
            }
            break
        }
        prev = d
        d = d.Next
    }
}


func (this *MyHashMap) GetBucket(key int, create bool) *HashMapTens {
    t := &this.HashMapTens
    k := key
    for k>10 {
        if k%10 == 0 {
            if t.bucket0 == nil {
                if create {
                    t.bucket0 = new(HashMapTens)
                } else {
                    break
                }
            }
            t = t.bucket0
        } else if k%10 == 1 {
            if t.bucket1 == nil {
                if create {
                    t.bucket1 = new(HashMapTens)
                } else {
                    break
                }
            }
            t = t.bucket1
        } else if k%10 == 2 {
            if t.bucket2 == nil {
                if create {
                    t.bucket2 = new(HashMapTens)
                } else {
                    break
                }
            }
            t = t.bucket2
        } else if k%10 == 3 {
            if t.bucket3 == nil {
                if create {
                    t.bucket3 = new(HashMapTens)
                } else {
                    break
                }
            }
            t = t.bucket3
        } else if k%10 == 4 {
            if t.bucket4 == nil {
                if create {
                    t.bucket4 = new(HashMapTens)
                } else {
                    break
                }
            }
            t = t.bucket4
        } else if k%10 == 5 {
            if t.bucket5 == nil {
                if create {
                    t.bucket5 = new(HashMapTens)
                } else {
                    break
                }
            }
            t = t.bucket5
        } else if k%10 == 6 {
            if t.bucket6 == nil {
                if create {
                    t.bucket6 = new(HashMapTens)
                } else {
                    break
                }
            }
            t = t.bucket6
        } else if k%10 == 7 {
            if t.bucket7 == nil {
                if create {
                    t.bucket7 = new(HashMapTens)
                } else {
                    break
                }
            }
            t = t.bucket7
        } else if k%10 == 8 {
            if t.bucket8 == nil {
                if create {
                    t.bucket8 = new(HashMapTens)
                } else {
                    break
                }
            }
            t = t.bucket8
        } else if k%10 == 9 {
            if t.bucket9 == nil {
                if create {
                    t.bucket9 = new(HashMapTens)
                } else {
                    break
                }
            }
            t = t.bucket9
        }
        k/=10
        if t.Next == nil {
            if create {
                t.Next = new(HashMapTens)
            } else {
                break
            }
        }
        t = t.Next
    }
    return t
}
