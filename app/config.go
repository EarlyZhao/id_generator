package app


type Config struct{
  Server *ServerConfig
  App *AppConfig
}


type ServerConfig struct{
  Addr string
  Port string
  Daemon bool
  LogPath string
  PidPath string
}


type AppConfig struct{

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
    App: NewAppConfig(),
  }
}

// must be excuted after App initilized
func WriteServerConfig(config map[string]interface{}){
  App.Config.Server.Addr = config["addr"].(string)
  App.Config.Server.Port = config["port"].(string)
  App.Config.Server.Daemon = config["daemon"].(bool)

  if App.Config.Server.Daemon{
    App.Config.Server.LogPath = config["log_path"].(string)
    App.Config.Server.PidPath = config["pid_path"].(string)
  }
}

func (c *Config) RunAddr() string{
  return c.Server.Addr + ":" + c.Server.Port
}