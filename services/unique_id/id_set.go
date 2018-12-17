package unique_id

import(
  "sync"
)

type IdSet struct{
  Set []*Buffer

  // Count defualt 1, is the max index of Set
  // len(Set) = Count + 1
  Count int

  Index int
  mutex sync.Mutex

}

func GetNewIdSet() *IdSet{
  set := &IdSet{Count: 1, Index: 0, Set: []*Buffer{} }
  for i := 0; i <= set.Count; i++{
    set.Set = append(set.Set, NewBuffer(i))
  }

  return set
}

func (i *IdSet) GetId() uint64{
  i.mutex.Lock()

  id, duration, remain := i.Current().ReleaseId()
  i.CheckBufferCondition(duration, remain)

  i.mutex.Unlock()
  return id
}

func (i *IdSet) CheckBufferCondition(duration uint64, remain uint64) {
  if ifSwitching(duration, remain) {
    // need a lock, not thread safe
    // Switch the Index of Set, to make Current() access next Buffer
    if i.Index == i.Count {
      i.Index = 0
    }else{
      i.Index += 1
    }
    // ensure buffer is not empty.
    // if the goroutine for getting buffer full, has not finished,
    // try GetBufferFull again
    i.Current().GetBufferFull()
  }
}

func (i *IdSet) Current() *Buffer{
  return i.Set[i.Index]
}
// define the timing of switching
func ifSwitching(duration uint64, remain uint64) bool{
  return remain == 0
}