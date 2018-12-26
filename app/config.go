package app


import(
  "github.com/id_generator/helpers"
  "fmt"
)


type Config struct{
  Server *ServerConfig
  Run *AppConfig
}


type ServerConfig struct{
  Addr string
  Port string
  Daemon bool
  LogPath string
  PidPath string
}

type DbConfig struct{
  Host string
  Port int
  Username string
  Password string
  Database string
}

type AppConfig struct{
  Secret string
  Database string
  Mysql *DbConfig
  Pg *DbConfig
}



func NewServerConfig() *ServerConfig{
  return &ServerConfig{}
}

func NewAppConfig() *AppConfig{
  return &AppConfig{}
}

func NewConfig() *Config{
  return &Config{
    Server: NewServerConfig(),
    Run: NewAppConfig(),
  }
}

// must be excuted after App initilized
func WriteConfig(config map[interface{}]interface{}){
  App.Config.Server.Addr = config["addr"].(string)
  App.Config.Server.Port = config["port"].(string)
  App.Config.Server.Daemon = config["daemon"].(bool)

  config_path := config["conf_path"].(string)
  var run AppConfig
  helpers.ReadConfig(config_path, &run) //.(*AppConfig)

  App.Config.Run = &run
  fmt.Println(App.Config.Run.Mysql)

  if App.Config.Server.Daemon{
    App.Config.Server.LogPath = config["log_path"].(string)
    App.Config.Server.PidPath = config["pid_path"].(string)
  }
}

func (c *Config) RunAddr() string{
  return c.Server.Addr + ":" + c.Server.Port
}

// func (c *AppConfig) GetDbConfig() (value map[interface{}]interface{}){
//   // value, _ = c.Value[key].(map[interface{}]interface{})
//   value["database"] = c.Database
//   return value
// }



