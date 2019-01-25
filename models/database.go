package models

import (
  "github.com/jinzhu/gorm"
  // _ "github.com/jinzhu/gorm/dialects/postgres"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "fmt"
  "github.com/id_generator/conf"
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
  <- conf.ConfigInitOverForDb

  var err error
  var dbUrl string
  config_db := conf.Config.Database

  if config_db == "mysql"{
    dbUrl = mysqlConnectionUrl()
    DB, err = gorm.Open("mysql", dbUrl)
  }else{
    // todo: pg
  }

  if err != nil{
    fmt.Println(conf.Config.Database)
    fmt.Println(conf.Config.Mysql)
    fmt.Println(dbUrl)
    panic(err)
  }
  // todo:
  DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&List{})
  // DB.Ping()

  ConnectionSucess <- true
}

func mysqlConnectionUrl() string{
  config := conf.Config

  user     := config.Mysql.Username
  password := config.Mysql.Password
  host     := config.Mysql.Host
  port     := config.Mysql.Port

  url      := fmt.Sprintf("%s:%d", host, port)
  database  := config.Mysql.Database

  dbUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, url, database)

  return dbUrl
}
