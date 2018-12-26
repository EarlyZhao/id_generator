package controllers

import(
  "net/http"
  "github.com/id_generator/context"
  "encoding/json"
  "fmt"
  "github.com/id_generator/sys_errors"
  "github.com/id_generator/helpers"
  "strconv"
)

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
  c.JsonResponse()
}

func (c *Controller) JsonResponse(){
  if jsonData, ok := c.Data["json"]; ok{
    if jsonString , err := json.Marshal(jsonData); err == nil{
      c.SetJsonBody(jsonString)
    }else{
      fmt.Println(err)
    }

  }
}

func (c * Controller) SetJsonBody(body []byte){
  c.Context.Output.SetBody(body)
  c.Context.Output.ContentJson()
}

func (c *Controller) RecoverFunc(context *context.Context){
  if err := recover(); err != nil{
    if err == sys_errors.ParamsError{
      return
    }

    panic(err)
  }
}

func (c *Controller) MustGetString(key string, msg string) string{
  value := c.Context.Input.GetString(key, "")
  if value != ""{
    return value
  }

  c.RaiseParamsError(msg)

  return ""
}

func (c *Controller) MustGetInt(key string, msg string) uint64{
  value := c.Context.Input.GetString(key, "")
  var ret uint64
  if value != ""{
    ret, err := strconv.ParseUint(value, 10, 64)
    if err == nil{
      return ret
    }
  }

  c.RaiseParamsError(msg)
  return ret
}

func (c *Controller)RaiseParamsError(msg string){
  data := helpers.WriteParamsErrorData(msg)
  c.SetJsonBody(data)

  panic(sys_errors.ParamsError)
}


