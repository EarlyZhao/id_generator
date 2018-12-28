package helpers

type SuccessRet struct{
  Status int `json: "status"`
  Data interface{} `json:"data"`
}

func NewSuccessRet(data interface{}) *SuccessRet{
  return &SuccessRet{Status: 1, Data: data,}
}