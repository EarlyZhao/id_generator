package logs

import(
  "sync"
  "os"
)

type ConsoleLogger struct{

}

var mutex sync.Mutex

func NewConsoleLogger() *ConsoleLogger{
  return &ConsoleLogger{}
}

func(c *ConsoleLogger) Name() string{
  return "Console"
}

func(c *ConsoleLogger) Close() {

}

func(c *ConsoleLogger) WriteMsg(msg *LogMsg){
  str := content(msg)
  writeToConsole([]byte(str))
}

func writeToConsole(bytes []byte){
  os.Stdout.Write(bytes)
}

func content(lm *LogMsg) string{
  return lm.Msg() + "\n"
}