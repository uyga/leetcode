package leetcode

type Tens struct {
    bucket0 *Tens
    bucket1 *Tens
    bucket2 *Tens
    bucket3 *Tens
    bucket4 *Tens
    bucket5 *Tens
    bucket6 *Tens
    bucket7 *Tens
    bucket8 *Tens
    bucket9 *Tens
    Next *Tens
    Data *ListNode
}

type MyHashSet struct {
    Tens
}

func MyHashSetConstructor() MyHashSet {
    return MyHashSet{}
}

func (this *MyHashSet) Add(key int)  {
    t := this.GetBucket(key, true)
    if t.Data == nil {
        t.Data = new(ListNode)
        t.Data.Val = key
    } else {
        d := t.Data
        for d.Next != nil && d.Val != key {
            d = d.Next
        }
        if d.Next == nil && d.Val != key {
            d.Next = new(ListNode)
            d.Next.Val = key
        }
    }
}

func (this *MyHashSet) Remove(key int)  {
    t := this.GetBucket(key, false)
    d := t.Data
    prev := d
    for d != nil {
        if d.Val == key {
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

func (this *MyHashSet) Contains(key int) bool {
    var res bool
    t := this.GetBucket(key, false)
    d := t.Data
    for d != nil {
        if d.Val == key {
            res = true
            break
        }
        d = d.Next
    }
    return res
}

func (this *MyHashSet) GetBucket(key int, create bool) *Tens {
    t := &this.Tens
    k := key
    for k>10 {
        if k%10 == 0 {
            if t.bucket0 == nil {
                if create {
                    t.bucket0 = new(Tens)
                } else {
                    break
                }
            }
            t = t.bucket0
        } else if k%10 == 1 {
            if t.bucket1 == nil {
                if create {
                    t.bucket1 = new(Tens)
                } else {
                    break
                }
            }
            t = t.bucket1
        } else if k%10 == 2 {
            if t.bucket2 == nil {
                if create {
                    t.bucket2 = new(Tens)
                } else {
                    break
                }
            }
            t = t.bucket2
        } else if k%10 == 3 {
            if t.bucket3 == nil {
                if create {
                    t.bucket3 = new(Tens)
                } else {
                    break
                }
            }
            t = t.bucket3
        } else if k%10 == 4 {
            if t.bucket4 == nil {
                if create {
                    t.bucket4 = new(Tens)
                } else {
                    break
                }
            }
            t = t.bucket4
        } else if k%10 == 5 {
            if t.bucket5 == nil {
                if create {
                    t.bucket5 = new(Tens)
                } else {
                    break
                }
            }
            t = t.bucket5
        } else if k%10 == 6 {
            if t.bucket6 == nil {
                if create {
                    t.bucket6 = new(Tens)
                } else {
                    break
                }
            }
            t = t.bucket6
        } else if k%10 == 7 {
            if t.bucket7 == nil {
                if create {
                    t.bucket7 = new(Tens)
                } else {
                    break
                }
            }
            t = t.bucket7
        } else if k%10 == 8 {
            if t.bucket8 == nil {
                if create {
                    t.bucket8 = new(Tens)
                } else {
                    break
                }
            }
            t = t.bucket8
        } else if k%10 == 9 {
            if t.bucket9 == nil {
                if create {
                    t.bucket9 = new(Tens)
                } else {
                    break
                }
            }
            t = t.bucket9
        }
        k/=10
        if t.Next == nil {
            if create {
                t.Next = new(Tens)
            } else {
                break
            }
        }
        t = t.Next
    }
    return t
}
