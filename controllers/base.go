package controllers

import(
  "net/http"
  "github.com/EarlyZhao/id_generator/context"
  "encoding/json"
  "fmt"
  "github.com/EarlyZhao/id_generator/sys_errors"
  "github.com/EarlyZhao/id_generator/helpers"
)

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

}

func (c *Controller) Done(){ // render
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
  value := c.Context.Input.GetUint(key, 0)
  var ret uint64
  if value != 0{
    return value
  }

  c.RaiseParamsError(msg)
  return ret
}

func (c *Controller) MustEqual(value interface{}, compare interface{}, msg string){
  if value == compare{ return }

  c.RaiseParamsError(msg)
}

func (c *Controller) RaiseParamsError(msg string){
  data := helpers.WriteParamsErrorData(msg)
  c.SetJsonBody(data)

  panic(sys_errors.ParamsError)
}


