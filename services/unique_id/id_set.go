package unique_id

import(
  "sync"
)

type IdSet struct{
  Set []*Buffer

  Count int // len(Set)
  SetIndexMax int // Count - 1

  Index int
  mutex sync.Mutex

}

func GetNewIdSet(count int, source interface{}) *IdSet{
  if count < 2{
    count = 2
  }
  set := &IdSet{Count: count, SetIndexMax: count - 1, Index: 0, Set: []*Buffer{} }

  set.fullSet(source)

  return set
}

func (i *IdSet) Reload(count int, source interface{}){
  i.mutex.Lock()
  defer i.mutex.Unlock()

  i.Count = count
  i.Set = []*Buffer{}

  i.fullSet(source)
}

func (i *IdSet) fullSet(source interface{}){
  for j := 0; j < i.Count; j++{
    i.Set = append(i.Set, NewBuffer(j, source))
  }
}

func (i *IdSet) GetId() uint64{
  i.mutex.Lock()
  defer i.mutex.Unlock()

  id, duration, remain := i.Current().ReleaseId()
  i.CheckBufferCondition(duration, remain)

  return id
}

func (i *IdSet) CheckBufferCondition(duration uint64, remain uint64) {
  if ifSwitching(duration, remain) {
    // need a lock, not thread safe
    // Switch the Index of Set, to make Current() access next Buffer
    if i.Index == i.SetIndexMax {
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