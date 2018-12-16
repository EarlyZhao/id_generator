package dispatcher

import(
  "net/http"
  "reflect"
  "github.com/id_generator/middlewares"
  "github.com/id_generator/context"
  "github.com/id_generator/controllers"
  // "fmt"
)
type RouteServe struct{
  // middlewares.MiddlewareBridge
  routes *Tree
}

type Context struct {
  Params map[string][]string
  Request *http.Request
}

func (t * RouteServe) MiddlewareCall(mr *middlewares.MiddlewareResponse, r *http.Request){
  method := r.Method
  path := r.URL.Path
  handler, params, controllerMethod ,err := t.routes.FindHandler(method + path)

  if err != nil{
    mr.Res404()
    return
  }

  //hr := &HandlerRequest{Params: r.URL.Query(), Request: r}
  httpParams := r.URL.Query()
  for key, value := range(params){
    httpParams[key] = []string{value}
  }

  context := context.NewContext()
  context.Input.Params = httpParams
  context.Input.Request = r
  context.Output.Response = mr

  var execController controllers.ControllerInterface
  // handlerType := reflect.ValueOf(handler).Elem().Type()
  handlerValue := reflect.ValueOf(handler)
  handlerType := reflect.Indirect(handlerValue).Type()

  handlerController := reflect.New(handlerType)
  execController = handlerController.Interface().(controllers.ControllerInterface)
  execController.Init(context)
    // code, headers, body := handler.Serve(httpParams, r)
  // execController.Serve(context)
  rv := reflect.ValueOf(execController)
  handleMethod := rv.MethodByName(controllerMethod)
  handleMethod.Call([]reflect.Value{})
  execController.Done()

}
