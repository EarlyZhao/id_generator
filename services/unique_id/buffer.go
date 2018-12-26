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
  Initialized bool
  mutex sync.Mutex // Lock sync.

  BufferSource interface{}
}

type BufferInterface interface{
  Duration() uint64
  StartAt() uint64
  EndAT() uint64
  Update()
}


func NewBuffer(index int, s interface{}) *Buffer{
  buffer := &Buffer{Num: index, Fulling: false, Initialized: false,}
  buffer.SetSource(s)
  buffer.GetBufferFull()
  return buffer
}


func (b *Buffer) SetSource( s interface{}){
  b.BufferSource = s
}

func (b *Buffer) ReleaseId() (id uint64, duration uint64, remaining uint64){
  // need a lock, or executed in a lock
  if b.Current < b.End{
    id = b.Current

    b.Current += 1  // not an atomic operation
    b.Cap = b.End - b.Current

    duration = b.Duration
    remaining =  b.Cap

  }else{ // Buffer was empty
    // CheckBufferCondition() will avoid this situation
  }

  if b.timeToFull(){
    go b.GetBufferFull()
  }else{ // buffer is not empty
    if b.Fulling{
      // when buffer began to ReleaseId(),
      // make it possible for goroutine to full the buffer later
      b.Fulling = false
    }
  }

  return id, duration, remaining
}

func (b *Buffer) GetBufferFull(){
  // redundancy check, reduce mutex grabbing as much as possible
  if b.Fulling{  // the buffer is full now
    return
  }

  b.mutex.Lock()
  if b.Fulling == false{

    if b.Initialized {
      b.Start = b.End
      b.End = b.End + b.Duration
    }else{
      b.Start = 10
      b.End = 20
      b.Duration = 10
      b.Initialized = true
    }

    b.Current = b.Start
    b.Cap = b.End - b.Current
  }
  // ensure just one goroutine to full buffer,
  // until the buffer began to ReleaseId().
  // it also represented the buffer is fulling over
  b.Fulling = true
  b.mutex.Unlock()
}

func (b *Buffer) timeToFull() bool{
  return b.Cap == 0
}
