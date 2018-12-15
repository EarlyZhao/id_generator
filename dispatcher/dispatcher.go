package dispatcher

import (
        "fmt"
        "net/http"
        "strings"
        "id_generator/middlewares"
        "reflect"
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

  path := r.URL.Path
  // h.Dispatch(path)
  fmt.Println(path)
  mR := middlewares.NewMiddlewareResponse()
  for _, middleware := range(h.Middleware){
    middleware.MiddlewareCall(mR, r)
  }

  mR.WriteResponse(rw)
}


func (h *Dispatcher) AddRoute(method string, path string, handleMethod string, handler interface{}) error {
  key := strings.ToUpper(method) + path
  h.routes.Add(key, handleMethod, handler)
  return nil
}


