package logs

import(
  "sync"
  "time"
)

// logger

type Logging struct{
  output Logger

  asynchronous bool
  mutex sync.Mutex

  msgChan chan *LogMsg
  signalChan chan string

}

type Logger interface{
  Name() string
  Close()
  WriteMsg(msg *LogMsg)
}

type LogMsg struct{
  level int
  msg string
  time time.Time
}

var Log *Logging
var msgPool *sync.Pool

var LogLevelDebug int = 0
var LogLevelInfo int = 1
var LogLevelWarn int = 2
var LogLevelError int = 3

var MsgChanLength int = 5000 // public parameter

func init(){
  msgPool = &sync.Pool{
    New: func() interface{} {
      return &LogMsg{}
    },
  }
}

func Info(msg string){
  if Log == nil{
    ConfigLogging(NewConsoleLogger())
  }

  Log.info(msg)
}

func (l *Logging) info(msg string){
  logMsg := msgPool.Get().(*LogMsg)

  logMsg.level = LogLevelInfo
  logMsg.msg = msg
  logMsg.time = time.Now()

  if l.asynchronous{
    l.msgChan <- logMsg
  }else{
    l.output.WriteMsg(logMsg) // output right now
    msgPool.Put(logMsg) // back to pool
  }
}

func ConfigLogging(output Logger){

  Log = &Logging{
    output: output, asynchronous: false,
  }

  // Log.Async() // todo: from config
}

func (l *Logging) Async(){
  l.mutex.Lock()
  defer l.mutex.Unlock()

  if l.asynchronous { return } // already started

  l.msgChan = make(chan *LogMsg, MsgChanLength)

  go l.startAsyncLogging()
  l.asynchronous = true
}

func (l *Logging) startAsyncLogging(){
  for{
    select{
      case logMsg := <- l.msgChan:
        l.output.WriteMsg(logMsg)
        msgPool.Put(logMsg)
    }
  }
}

func (lm *LogMsg) Level() int{
  return lm.level
}

func (lm *LogMsg) Msg() string{
  return lm.msg
}

func (lm *LogMsg) Time() time.Time{
  return lm.time
}
