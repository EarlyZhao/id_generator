package helpers

import(
  "io/ioutil"
  "gopkg.in/yaml.v2"
  "fmt"
)


func ReadConfig(path string, t interface{}) interface{}{
  data, err := ioutil.ReadFile(path)
  if err != nil{
    panic(fmt.Sprintf("Read Config file error: %v",err))
  }
  
  err = yaml.Unmarshal(data, t)
  if err != nil{
  	panic(err)
  }
  return t
}