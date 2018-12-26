package models

import (
  "github.com/jinzhu/gorm"
  // _ "github.com/jinzhu/gorm/dialects/postgres"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "fmt"
  "github.com/id_generator/app"
)

var DB *gorm.DB
var ConnectionSucess chan bool

func init(){
  ConnectionSucess = make(chan bool)
  go connectionToDB()
}

func connectionToDB(){
  // wait for app filished init process
  // database connection need the config data
  <- app.App.InitOver

  var err error
  config_db := app.App.Config.Run.Database
  fmt.Println(config_db)
  if config_db == "mysql"{
    DB, err = gorm.Open("mysql", mysqlConnectionUrl())
  }else{
    // todo: pg
  }

  if err != nil{
    panic(err)
  }
  // todo:
  DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&List{})
  // DB.Ping()

  ConnectionSucess <- true
}

func mysqlConnectionUrl() string{
  config := app.App.Config.Run

  user     := config.Mysql.Username
  password := config.Mysql.Password
  host     := config.Mysql.Host
  port     := config.Mysql.Port

  url      := fmt.Sprintf("%s:%d", host, port)
  database  := config.Mysql.Database

  dbUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, url, database)
  fmt.Println(dbUrl)
  return dbUrl
}
