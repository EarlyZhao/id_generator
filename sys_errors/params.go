package sys_errors


import(
  "encoding/json"
)

var ParamsError error

type SysError struct{
  ErrorCode int `json:"error_code"`
  Msg string   `json:"msg"`
}

func NewSysError(msg string) *SysError{
  return &SysError{ErrorCode: 1001, Msg: msg,}
}

func (s *SysError) JsonString() ([]byte, error){
  var err error
  if jsonString, err := json.Marshal(s); err == nil{
    return jsonString, nil
  }

  return []byte(""), err
}