package middlewares

import(
  "net/http"
  )
type Middleware interface{
  MiddlewareCall(rw *MiddlewareResponse, r *http.Request)
}

type MiddlewareBridge struct{
  // Bridge func(rw http.ResponseWriter, r *http.Request)
  Bridge Middleware
}

type MiddlewareResponse struct {
  Code int
  Headers map[string]string
  Body []byte
  Err error
}

func (mr *MiddlewareResponse) Header(key string, value string){
  mr.Headers[key] = value
}

func (mr *MiddlewareResponse) Res404(){
  mr.Code = 404
}

func (mr *MiddlewareResponse) Res500(msg string){
  mr.Code = 500
  mr.SetBody([]byte(msg))
}

func (mr *MiddlewareResponse) WriteError(err error){
  mr.Err = err
}

func (mr *MiddlewareResponse) Error() string{
  if mr.Err ==nil{
    return ""
  }

  return mr.Err.Error()
}

func (mr *MiddlewareResponse) WriteResponse(rw http.ResponseWriter){
  for key , value := range(mr.Headers){
    rw.Header().Add(key, value)
  }

  rw.WriteHeader(mr.Code)

  if len(mr.Body) > 0{
    rw.Write(mr.Body)
  }
}

func (mr *MiddlewareResponse) SetBody(data []byte){
  mr.Body = data
}

func NewMiddlewareResponse() *MiddlewareResponse{
  mr := MiddlewareResponse{Code: 200, Headers: make(map[string]string), Body: []byte(""),}
  return &mr
}

func DefaultMiddlewares() []Middleware{
  var middlewares []Middleware
  middlewares = append(middlewares, &ServeTime{})

  return middlewares
}