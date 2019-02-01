package conf

import(
  "github.com/id_generator/helpers"
)
var Config AppConfig
var ConfigInitOverForDb chan bool
var ConfigInit bool
var Run *RunConfig

func init(){
  ConfigInitOverForDb = make(chan bool, 1) // no blocking for ReadConfig
  Run = &RunConfig{Logging: true}
}


type AppConfig struct{
  Secret string
  Buffers int // the count of buffers
  Database string
  Mysql *DbConfig
  Pg *DbConfig
}

type DbConfig struct{
  Host string
  Port int
  Username string
  Password string
  Database string
}

type RunConfig struct{
  Logging bool
  Mode int
}

func ReadConfig(path string){
  if ConfigInit{ return }
  helpers.ReadConfig(path, &Config) //.(*AppConfig)

  setDefault()

  ConfigInit = true
  ConfigInitOverForDb <- true
}

func SetLogging(f bool){
  Run.Logging = f
}

func Logging() bool{
  return Run.Logging
}

func setDefault(){
  if Config.Buffers < 2{
    Config.Buffers = 2
  }
}