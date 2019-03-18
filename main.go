package main


import(
        "github.com/EarlyZhao/id_generator/app"
        _ "github.com/EarlyZhao/id_generator/routes"
        "flag"
        "fmt"
        "runtime"
      )

func main() {

  var addr        = flag.String("addr", "127.0.0.1", "the IP listen to")
  var port        = flag.String("port", "1314", "The Port bind on for HTTP or gRPC, you can pass a blank value to not use it")
  var close_log   = flag.Bool("close_log", false, "Whether to disable logging")
  var config_path = flag.String("config_path", "", "The config file path")
  var max_procs   = flag.Int("max_procs", 0, "The max count of goroutings that parallel running")
  var run         = flag.String("run", "http", "Http|grpc, the server run as HTTP or gRPC")
  flag.Parse()

  config := make(app.ConfigType)

  config["addr"] = *addr
  config["port"] = *port
  config["conf_path"] = *config_path
  config["run"] = *run
  config["logging"] = !*close_log

  if *config_path == ""{
    panic("config_path must be present")
  }

  var max_count_procs int
  max_count_procs = *max_procs
  if max_count_procs == 0{
    max_count_procs = runtime.NumCPU()
    if max_count_procs > 1 {
      max_count_procs = max_count_procs/2
    }
    fmt.Printf("set max_procs default as: %d\n", max_count_procs)
  }
  config["max_procs"] = max_count_procs
  runtime.GOMAXPROCS(max_count_procs)

  fmt.Println("Server Config:")
  for key, value := range(config){
    fmt.Println(key, ":", value)
  }

  app.Ready(config)
  app.Run()
}