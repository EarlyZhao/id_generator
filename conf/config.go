package conf

import(
  "github.com/id_generator/helpers"
)
var Config AppConfig
var ConfigInitOverForDb chan bool
var ConfigInit bool

func init(){
  ConfigInitOverForDb = make(chan bool, 1) // no blocking for ReadConfig
}


type AppConfig struct{
  Secret string
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

func ReadConfig(path string){
  if ConfigInit{ return }
  helpers.ReadConfig(path, &Config) //.(*AppConfig)

  ConfigInit = true
  ConfigInitOverForDb <- true
}