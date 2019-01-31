package test

import(
  "testing"
  // "unsafe"
  "fmt"
  "runtime"
)

func MustEqual(t *testing.T , result interface{}, compare interface{}, whenFailed string, whenSuccess string){
  if result != compare{
    printCaller()
    t.Logf("result: %v, %T, -- compare: %v, %T", result, result, compare, compare)
    t.Error(whenFailed)
  }else if whenSuccess != ""{
    t.Log(whenSuccess)
  }
}

func MustNoEqual(t *testing.T, result interface{}, compare interface{}, whenFailed string, whenSuccess string){
  if result == compare{
    printCaller()
    t.Error(whenFailed)
  }else if whenSuccess != ""{
    t.Log(whenSuccess)
  }
}

func MustGreaterUint(t *testing.T, bigger uint64, compare uint64, whenFailed string, whenSuccess string){
  // biggerUint := *(*uint64)(unsafe.Pointer(&bigger))
  // compareUint := *(*uint64)(unsafe.Pointer(&compare))
  if bigger <= compare{
    printCaller()
    t.Logf("greater: %v, compare: %v", bigger, compare)
    t.Error(whenFailed)
  }else if whenSuccess != ""{
    t.Log(whenSuccess)
  }
}

func printCaller(){
  var stack string
  for i := 1; ; i++ {
    _, file, line, ok := runtime.Caller(i)
    if !ok {
      break
    }

    stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))
  }
  fmt.Println(stack)
}