package app

import(
  "net/http"
  "net"
  "github.com/id_generator/dispatcher"
  grpc "google.golang.org/grpc"
  "github.com/id_generator/services/unique_id"
  "github.com/id_generator/grpc/id_rpc"
  "github.com/id_generator/logs"
  "fmt"
  "strings"
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

  if strings.ToUpper(App.Config.Server.RunAs) == "GRPC"{
    lis, err := net.Listen("tcp", App.Config.RunAddr())
    if err != nil {
      panic(fmt.Sprintf("gRPC failed to listen: %v", err))
    }

    s := grpc.NewServer()
    id_rpc.RegisterUniqueIdServiceServer(s, &unique_id.UniqueIdRpcService{})

    logs.Info("gRPC Listening on "  + App.Config.RunAddr())
    s.Serve(lis)
  }else{
    logs.Info("HTTP Listening on " + App.Config.RunAddr())
    App.Server.ListenAndServe()
  }

}

func Ready(config map[interface{}]interface{}){
  WriteConfig(config)
  App.Dispatcher.Ready()
  App.Server.Addr = App.Config.RunAddr()
}

func AddRoute(method string, path string, handleMethod string, handler interface{}) error {
  return App.Dispatcher.AddRoute(method, path, handleMethod, handler)
}