package unique_id


import(
  "sync"
  "fmt"
  "time"
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

  BufferSource BufferInterface
}

type BufferInterface interface{
  Duration() uint64
  StartAt() uint64
  EndAT() uint64
  Update() error
  Usable() bool
}


func NewBuffer(index int, s interface{}) *Buffer{
  buffer := &Buffer{Num: index, Fulling: false, Initialized: false,}
  buffer.SetSource(s)
  buffer.GetBufferFull()
  return buffer
}


func (b *Buffer) SetSource( s interface{}){
  b.BufferSource = s.(BufferInterface)
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
  defer b.mutex.Unlock()

  if b.Fulling == false{
    if err := b.BufferSource.Update(); err != nil{
      errMsg := fmt.Sprintf("GetBufferFull Error for: %v", err)
      fmt.Println(errMsg)
      // if current method running in a gorouting, the panic will make the process crash
      panic(err)
    }

    b.Start = b.BufferSource.StartAt()
    b.Duration = b.BufferSource.Duration()

    b.End = b.BufferSource.EndAT()
    b.Current = b.End - b.Duration
    b.Initialized = true
    b.Cap = b.End - b.Current
    fmt.Println("GetBufferFull At:", time.Now().Format("2006-01-02 15:04:05"), b.Current, b.End)

    // ensure just one goroutine to full buffer,
    // until the buffer began to ReleaseId().
    // it also represented the buffer is fulling over
    b.Fulling = true
  }
}

func (b *Buffer) timeToFull() bool{
  return b.Cap == 0
}
