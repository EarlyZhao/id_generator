package dispatcher

import (
        "fmt"
        "net/http"
        "strings"
        "github.com/id_generator/middlewares"
        "reflect"
        "time"
        "runtime"
       )


type Dispatcher struct{
  routes *Tree
  Middleware []middlewares.Middleware
}



// nesting middlewares
func (h *Dispatcher) Ready(){
  var routeMiddleware = []middlewares.Middleware{&RouteServe{routes: h.routes}}
  h.Middleware = append(routeMiddleware, h.Middleware...)

  var outMiddleware middlewares.Middleware
  var inMiddlware middlewares.Middleware

  for _, middleware := range(h.Middleware){
    if inMiddlware == nil{
      inMiddlware = middleware
    }else{
      fmt.Printf("(%v, %T)\n", middleware, middleware)
      middlewareElem := reflect.ValueOf(middleware).Elem()
      bridgeField := middlewareElem.FieldByName("Bridge")
      bridgeField.Set(reflect.ValueOf(inMiddlware))
      // elemField := execElem.FieldByName(fieldType.Name)
      // middleware.(reflect.TypeOf(middleware)).Bridge = inMiddlware
      inMiddlware = middleware
    }

  }

  outMiddleware = inMiddlware
  h.Middleware = []middlewares.Middleware{outMiddleware}
}

func NewDispatcher() *Dispatcher{
  return &Dispatcher{
      routes: &Tree{
        Root: &Node{ Children: make(map[string]*Node)},
      },
      Middleware: middlewares.DefaultMiddlewares(),
  }
}

func (h *Dispatcher) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
  start := time.Now()
  defer captureException(rw, r) // deal the errors that unexpected, prevention for crashing

  mR := middlewares.NewMiddlewareResponse()
  for _, middleware := range(h.Middleware){
    middleware.MiddlewareCall(mR, r)
  }

  mR.WriteResponse(rw)
  logging(mR, r, start)
}


func (h *Dispatcher) AddRoute(method string, path string, handleMethod string, handler interface{}) error {
  key := strings.ToUpper(method) + path
  h.routes.Add(key, handleMethod, handler)
  return nil
}

func captureException(rw http.ResponseWriter, r *http.Request){
  if err := recover(); err != nil{
    loggingException(r, err)

    rw.WriteHeader(500)
    rw.Write([]byte(fmt.Sprintln("%v",err)))
  }
}

func loggingException(r *http.Request, err interface{}){
  if err == nil{
    return
  }
  err_request := fmt.Sprintf("Request Error: %s %s %s : %s", r.Method, r.URL, time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf("%v",err))
  fmt.Println(err_request)
  var stack string
  for i := 1; ; i++ {
    _, file, line, ok := runtime.Caller(i)
    if !ok {
      break
    }

    stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))
  }
  fmt.Println(stack)
}

func logging(mr *middlewares.MiddlewareResponse, r *http.Request, start time.Time){
  str := fmt.Sprintf("%s %s %s %d %s  %s", r.Method, r.URL, time.Now().Format("2006-01-02~15:04:05"), mr.Code, time.Now().Sub(start).String(), mr.Error())
  fmt.Println(str)
}

