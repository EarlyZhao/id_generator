package main


import(
        "github.com/id_generator/app"
        _ "github.com/id_generator/routes"
        "flag"
        "fmt"
        "runtime"
      )

func main() {

  var addr        = flag.String("addr", "127.0.0.1", "the IP listen to")
  var port        = flag.String("port", "1314", "the Port bind on")
  var log_path    = flag.String("log_path", "/var/log/id_generator.log", "the log file path")
  var pid_path    = flag.String("pid_path", "/tmp/pids/id_generator.pid", "the pid file path")
  var config_path = flag.String("config_path", "", "the config file path")
  var daemon      = flag.Bool("daemon", false, "run as daemon")
  var max_procs   = flag.Int("max_procs", 0, "the max count of goroutings that parallel running")
  flag.Parse()

  config := make(map[interface{}]interface{})

  config["addr"] = *addr
  config["port"] = *port
  config["daemon"] = *daemon
  config["conf_path"] = *config_path

  if *config_path == ""{
    panic("conf_path must be present")
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

  if config["daemon"] == true{
    config["log_path"] = *log_path
    config["pid_path"] = *pid_path
  }

  fmt.Println("Server Config:")
  for key, value := range(config){
    fmt.Println(key, ":", value)
  }

  app.Ready(config)
  app.Run()
}