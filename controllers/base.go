package controllers

import(
  "net/http"
  "id_generator/context"
  "encoding/json"
  "fmt"
)

type ControllerInterface interface{
  Init(context *context.Context)
  Serve()
  Done()
}
// func (h *BaseHandler) AddHeader(key string, value interface{}){
//  h.RequestHeaders[key] = value
// }
type Controller struct{
  Params map[string][]string
  Request http.Request
  Data map[interface{}]interface{}
  Context *context.Context
}


func (c *Controller) Init(context *context.Context){
  c.Context = context
  c.Data = make(map[interface{}]interface{})
}

func (c *Controller) Serve(){
  // c.Context.Output.Response.Header("haha", "Nnm")
  // c.Context.Output.Response.Body = []byte(context.Input.Params["id"][0])
}

func (c *Controller) Done(){ // render
  // c.Context.Output.Response.Body = []byte(c.Data["json"].(string))
  if jsonData, ok := c.Data["json"]; ok{
    if jsonString , err := json.Marshal(jsonData); err == nil{
      c.Context.Output.SetBody(jsonString)
      c.Context.Output.ContentJson()
    }else{
      fmt.Println(err)
    }

  }
}