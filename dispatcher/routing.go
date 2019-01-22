package dispatcher

import(
  "net/http"
  "reflect"
  "github.com/id_generator/middlewares"
  "github.com/id_generator/context"
  "strings"
  "io/ioutil"
  "encoding/json"
)
type RouteServe struct{
  // middlewares.MiddlewareBridge
  routes *Tree
}

type ControllerInterface interface{
  Init(context *context.Context)
  Serve()
  Done()
  RecoverFunc(context *context.Context)
}

func (t * RouteServe) MiddlewareCall(mr *middlewares.MiddlewareResponse, r *http.Request){
  method := r.Method
  path := r.URL.Path

  handler, params, controllerMethod ,err := t.routes.FindHandler(method + path)

  if err != nil{
    mr.Res404()
    return
  }

  httpParams := make(context.InputParams)

  for key, value := range(r.URL.Query()){
    httpParams[key] = value
  }
  for key, value := range(params){
    httpParams[key] = value
  }

  method = strings.ToUpper(method) // ensure as upper
  // the value from Body has higher priority than URL query
  if method == "POST" || method == "PUT" || method == "PATCH"{
    body, err := ioutil.ReadAll(r.Body)
    if err != nil{}
    bodyData := make(context.InputParams)
    jsonErr := json.Unmarshal(body, &bodyData)
    if jsonErr == nil{
      for key_post, value_post := range(bodyData){
        httpParams[key_post] = value_post
      }
    }else{
      mr.WriteError(jsonErr)
      mr.Res500(jsonErr.Error())
      return
    }
  }

  context := context.NewContext()
  context.Input.Params = httpParams
  context.Input.Request = r
  context.Output.Response = mr

  var execController ControllerInterface
  handlerValue := reflect.ValueOf(handler)
  handlerType := reflect.Indirect(handlerValue).Type()

  handlerController := reflect.New(handlerType)
  execController = handlerController.Interface().(ControllerInterface)
  execController.Init(context)

  rv := reflect.ValueOf(execController)
  handleMethod := rv.MethodByName(controllerMethod)

  defer execController.RecoverFunc(context)
  handleMethod.Call([]reflect.Value{})
  execController.Done()

}
