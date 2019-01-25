package app


import(
  "github.com/id_generator/logs"
  "github.com/id_generator/conf"
)


type Config struct{
  Server *ServerConfig
}


type ServerConfig struct{
  Addr string
  Port string
  RunAs string
  Daemon bool
  LogPath string
  PidPath string
}

type ConfigType map[interface{}]interface{}

func NewServerConfig() *ServerConfig{
  return &ServerConfig{}
}

func NewLogger(config ConfigType){
  logger := logs.NewConsoleLogger() // todo: from config file
  logs.ConfigLogging(logger)
}

func NewConfig() *Config{
  return &Config{
    Server: NewServerConfig(),
  }
}

// must be excuted after App initilized
func WriteConfig(config ConfigType){
  App.Config.Server.Addr = config["addr"].(string)
  App.Config.Server.Port = config["port"].(string)
  App.Config.Server.Daemon = config["daemon"].(bool)
  App.Config.Server.RunAs = config["run"].(string)
  config_path := config["conf_path"].(string)

  conf.ReadConfig(config_path)

  if App.Config.Server.Daemon{ // todo: delete it
    App.Config.Server.LogPath = config["log_path"].(string)
    App.Config.Server.PidPath = config["pid_path"].(string)
  }

  NewLogger(config)
}

func (c *Config) RunAddr() string{
  return c.Server.Addr + ":" + c.Server.Port
}




