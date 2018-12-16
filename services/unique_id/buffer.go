package unique_id


import(
  // "sync"
)

type Buffer struct{
  Start int // the start value of id
  Current int // current id value, plus one per time when id was acquired
  End int // the max value of id

  Duration int // End - Start
  Cap int // End - Current

  Num int // the serial number
  Fulling bool // if getting the buffer full
  // Lock sync.
}


func NewBuffer(index int) *Buffer{
  buffer := &Buffer{Num: index, Fulling: false, }
  buffer.GetBufferFull()
  return buffer
}

func (b *Buffer) ReleaseId() (id int, duration int, remaining int){
  // need a lock
  // var id, duration, remaining int
  if b.Current < b.End{
    id = b.Current

    b.Current += 1
    b.Cap = b.End - b.Current

    duration = b.Duration
    remaining =  b.Cap

  }
  return id, duration, remaining
}

func (b *Buffer) GetBufferFull(){
  // for test
  b.Start = 10000
  b.End = 20000
  b.Current = b.Start
  b.Duration = 10000
  b.Cap = b.End - b.Current
}
