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
  InitOver chan bool
}

func init(){
  App = NewApp()
}

func NewApp() *Application{
  app := &Application{
            Server: &http.Server{},
            Config: NewConfig(),
            Dispatcher: dispatcher.NewDispatcher(),
            InitOver: make(chan bool),
          }
  app.Server.Handler = app.Dispatcher
  return app
}

func Run(){
  App.Server.ListenAndServe()
}

func Ready(config map[interface{}]interface{}){
  WriteConfig(config)
  App.Dispatcher.Ready()
  App.Server.Addr = App.Config.RunAddr()

  // the goroutings that wait for App filished init process, stop blocking now
  App.InitOver <- true
}

func AddRoute(method string, path string, handleMethod string, handler interface{}) error {
  return App.Dispatcher.AddRoute(method, path, handleMethod, handler)
}