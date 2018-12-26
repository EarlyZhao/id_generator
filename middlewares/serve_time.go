package middlewares

// import "github.com/id_generator/middlewate"

import(
      "time"
      "net/http"
    )

type ServeTime struct{
  MiddlewareBridge
}

func (s * ServeTime) MiddlewareCall(rw *MiddlewareResponse, r *http.Request){
  start := time.Now()

  if s.Bridge != nil{
    s.Bridge.MiddlewareCall(rw, r)
  }

  now := time.Now()
  duration := now.Sub(start)

  rw.Header("ServeTime", duration.String())
}