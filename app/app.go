package app

import(
  "net/http"
  "github.com/id_generator/dispatcher"
)

var App *Application


type Application struct{
  Server *http.Server
  Config *Config

  Dispatcher *dispatcher.Dispatcher
  // Middleware []*Middleware
}

func init(){
  App = NewApp()
}

func NewApp() *Application{
  app := &Application{
            Server: &http.Server{},
            Config: NewConfig(),
            Dispatcher: dispatcher.NewDispatcher(),
          }
  app.Server.Handler = app.Dispatcher
  return app
}

func Run(){
  App.Dispatcher.Ready()
  App.Server.Addr = App.Config.RunAddr()
  App.Server.ListenAndServe()
}

func AddRoute(method string, path string, handleMethod string, handler interface{}) error {
  return App.Dispatcher.AddRoute(method, path, handleMethod, handler)
}