package unique_id

import(
  "testing"
  "github.com/EarlyZhao/id_generator/test"
)

type TestList struct{
  duration uint64
  startAt uint64
  endAt uint64
  usable bool
  BusinessType string
}

func (l *TestList) Duration() uint64 { return l.duration }
func (l *TestList) StartAt()  uint64 { return l.startAt}
func (l *TestList) EndAT()    uint64 { return l.endAt}
func (l *TestList) Usable()   bool   { return l.usable}
func (l *TestList) Update()   error{
  l.endAt += l.duration
  return nil
}

func TestWareHouse(t *testing.T){
  var h *WareHouse = NewWareHouse()
  var lists []*TestList
  var testBusinessType string = "testing"
  var idSetLen int = 2
  var err error
  var id, startAt uint64 = 0, 100

  lists = append(lists, &TestList{duration: 1000, startAt: startAt, endAt: 1000, usable: true, BusinessType: testBusinessType})

  for _, list := range(lists){
    set := GetNewIdSet(idSetLen, list)
    h.AddNewToWareHouse(list.BusinessType, set)
  }

  _, ok := h.HouseMap[testBusinessType]
  test.MustEqual(t, ok, true, "InitWareHouse failed", "Pass Init WareHouse")

  id, err = h.Acquire(testBusinessType)
  test.MustEqual(t, err, nil, "WareHouse Acquire() Id failed", "")
  test.MustGreaterUint(t, id, startAt, "WareHouse Acquire() ID failed","")

  h.RemoveToWareHouse(testBusinessType)
  _, err = h.Acquire(testBusinessType)
  test.MustNoEqual(t, err, nil, "WareHouse RemoveToWareHouse() failed", "")
}

func TestIdSet(t *testing.T){
  var list *TestList
  var testBusinessType string = "testing"
  var idSetLen, resetLen int = 5, 3
  var set *IdSet
  var id, startAt uint64 = 0, 100

  list = &TestList{duration: 1000, startAt: startAt, endAt: 1000, usable: true, BusinessType: testBusinessType}
  set = GetNewIdSet(idSetLen, list)
  test.MustEqual(t, len(set.Set), idSetLen, "IdSet GetNewIdSet() failed", "")

  set.Reload(resetLen, list)
  test.MustEqual(t, len(set.Set), resetLen, "IdSet Reload() failed", "")

  index := set.Index
  current := set.Current()

  set.turnToNextBuffer()
  newCurrent := set.Current()

  test.MustNoEqual(t, index, set.Index, "IdSet turnToNextBuffer() failed", "")
  test.MustNoEqual(t, current, newCurrent, "IdSet turnToNextBuffer() failed", "")

  id = set.GetId()
  test.MustGreaterUint(t, id, startAt, "IdSet GetId() failed","")
}


func TestBuffer(t *testing.T){
  var buffer *Buffer
  var initDuration uint64 = 1000
  var id uint64
  var empty bool
  list := &TestList{duration: initDuration, startAt: 0, endAt: 1000, usable: true, BusinessType: "test"}

  buffer = NewBuffer(0, list)
  test.MustEqual(t, buffer.BufferSource, list, "Buffer NewBuffer() failed","")
  test.MustEqual(t, buffer.Fulling, true, "Buffer NewBuffer() Fulling failed","")

  id, empty = buffer.ReleaseId()
  test.MustEqual(t, id, initDuration, "Buffer ReleaseId() failed","")
  test.MustEqual(t, buffer.Duration, initDuration, "Buffer ReleaseId() failed when check duration","")
  test.MustEqual(t, buffer.Cap, initDuration - 1, "Buffer ReleaseId() failed when check cap","")
  test.MustEqual(t, empty, false, "Buffer ReleaseId() failed when check empty","")

  buffer.ReleaseId()
  test.MustGreaterUint(t, buffer.Current + buffer.Duration, buffer.End, "Buffer ReleaseId() wrong", "")

  buffer.Current = buffer.End - 1
  id, empty = buffer.ReleaseId()
  test.MustEqual(t, empty, true, "Buffer ReleaseId() failed when check empty","")

  buffer.GetBufferFull()
  test.MustEqual(t, buffer.Current + buffer.Duration, buffer.End, "Buffer GetBufferFull() failed", "")
}

