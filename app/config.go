package app


import(
  "github.com/id_generator/helpers"
  "fmt"
  "github.com/id_generator/logs"
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

type ConfigType map[interface{}]interface{}

func NewServerConfig() *ServerConfig{
  return &ServerConfig{}
}

func NewAppConfig() *AppConfig{
  return &AppConfig{}
}

func NewLogger(config ConfigType){
  logger := logs.NewConsoleLogger() // todo: from config file
  logs.ConfigLogging(logger)
}

func NewConfig() *Config{
  return &Config{
    Server: NewServerConfig(),
    Run: NewAppConfig(),
  }
}

// must be excuted after App initilized
func WriteConfig(config ConfigType){
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

  NewLogger(config)
}

func (c *Config) RunAddr() string{
  return c.Server.Addr + ":" + c.Server.Port
}




