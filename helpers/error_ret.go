package helpers


import(
  "github.com/id_generator/sys_errors"
)
type ErrorRet struct{
  Status int `json:"status"`
  Error interface{} `json:"error"`
}

func NewErrorRet(code int, msg string) *ErrorRet{
  err := sys_errors.NewSysError(code, msg)
  return &ErrorRet{Status: 0, Error: err, }
}