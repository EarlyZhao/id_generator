package main


import(
        "github.com/id_generator/app"
        _ "github.com/id_generator/routes"
        "flag"
        "fmt"
      )

func main() {
  var addr = flag.String("addr", "127.0.0.1", "the IP listen to")
  var port = flag.String("port", "1314", "the Port bind on")
  var log_path = flag.String("log_path", "/var/log/id_generator.log", "the log file path")
  var pid_path = flag.String("pid_path", "/tmp/pids/id_generator.pid", "the pid file path")
  var daemon = flag.Bool("daemon", false, "run as daemon")

  flag.Parse()

  config := make(map[string]interface{})

  config["addr"] = *addr
  config["port"] = *port
  config["daemon"] = *daemon

  if config["daemon"] == true{
    config["log_path"] = *log_path
    config["pid_path"] = *pid_path
  }

  fmt.Println("Server Config")
  for key, value := range(config){
    fmt.Println(key, ":", value)
  }
  fmt.Println("----------")

  app.WriteServerConfig(config)
  app.Run()
}