package sys_errors

import(
  "encoding/json"
)

type SysError struct{
  ErrorCode int `json:"error_code"`
  Msg string   `json:"msg"`
}


func (s *SysError) JsonString() ([]byte, error){
  var err error
  if jsonString, err := json.Marshal(s); err == nil{
    return jsonString, nil
  }

  return []byte(""), err
}

func (s *SysError) JsonMap() ([]byte, error){
  var err error
  if jsonString, err := json.Marshal(s); err == nil{
    return jsonString, nil
  }

  return []byte(""), err
}

func NewSysError(code int, msg string) *SysError{
  return &SysError{ErrorCode: code, Msg: msg,}
}
