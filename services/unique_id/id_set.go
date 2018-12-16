package unique_id

type IdSet struct{
  Set []*Buffer
  Count int  // defualt 2
  Index int
  // Current *Buffer
}


func GetNewIdSet() *IdSet{
  set := &IdSet{Count: 2, Index: 0, Set: []*Buffer{} }
  for i := 0; i < set.Count; i++{
    set.Set = append(set.Set, NewBuffer(i))
  }

  return set
}

func (i *IdSet) GetId() int{
  id, duration, remain := i.Current().ReleaseId()

  i.CheckBufferCondition(duration, remain)

  return id
}

func (i *IdSet) CheckBufferCondition(duration int, remain int) {

  if remain == 0 {
    // need a lock
    current := i.Current()
    if i.Index == i.Count - 1 {
      i.Index = 0
    }else{
      i.Index += 1
    }

    go current.GetBufferFull()
  }
}

func (i *IdSet) Current() *Buffer{
  return i.Set[i.Index]
}