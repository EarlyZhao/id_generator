package unique_id


import(
  "sync"
)

type Buffer struct{
  Start uint64 // the start value of id
  Current uint64 // current id value, plus one per time when id was acquired
  End uint64 // the max value of id

  Duration uint64 // End - Start
  Cap uint64 // End - Current

  Num int // the serial number
  Fulling bool // if getting the buffer full
  mutex sync.Mutex // Lock sync.
}


func NewBuffer(index int) *Buffer{
  buffer := &Buffer{Num: index, Fulling: false, }
  buffer.GetBufferFull()
  return buffer
}

func (b *Buffer) ReleaseId() (id uint64, duration uint64, remaining uint64){
  // need a lock
  // var id, duration, remaining int
  if b.Current < b.End{
    id = b.Current

    b.Current += 1  // not an atomic operation
    b.Cap = b.End - b.Current

    duration = b.Duration
    remaining =  b.Cap

  }

  if b.timeToFull(){
    go b.GetBufferFull()
  }else{ // buffer is not empty
    if b.Fulling{
      b.Fulling = false
    }
  }

  return id, duration, remaining
}

func (b *Buffer) GetBufferFull(){
  // for test
  b.mutex.Lock()
  if b.Fulling == false{
    b.Fulling = true

    b.Start = 10000
    b.End = 20000
    b.Current = b.Start
    b.Duration = 10000
    b.Cap = b.End - b.Current

    // b.Fulling = false
  }
  b.mutex.Unlock()
}

func (b *Buffer) timeToFull() bool{
  return b.Cap == 0
}
