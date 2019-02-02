package context

import(
  "github.com/EarlyZhao/id_generator/middlewares"
)

const (
  ApplicationJSON = "application/json; charset=utf-8"
  ApplicationXML  = "application/xml; charset=utf-8"
  ApplicationYAML = "application/x-yaml; charset=utf-8"
  TextXML         = "text/xml; charset=utf-8"
)

type Output struct{
  Response *middlewares.MiddlewareResponse
}

func (o *Output) SetHeader(key string, value string){
  o.Response.Header(key, value)
}

func (o *Output) ContentJson(){
  o.SetHeader("Content-Type", ApplicationJSON)
}

func (o *Output) SetBody(value []byte){
  o.Response.SetBody(value)
}

func (o *Output) SetCode(code int){
  o.Response.Code = code
}