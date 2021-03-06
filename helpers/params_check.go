package helpers


import(
  "github.com/EarlyZhao/id_generator/sys_errors"
  "fmt"
)

type ParamsError struct{

}

func WriteParamsErrorData(msg string) []byte{
  sysError := sys_errors.NewSysError(1001, msg)
  data, err := sysError.JsonString()
  if err == nil{
    return data
  }

  // logging the error
  fmt.Println(err)
  return []byte("")
}