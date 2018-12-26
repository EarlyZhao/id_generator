package helpers

import(
  "io/ioutil"
  "gopkg.in/yaml.v2"
)


func ReadConfig(path string, t interface{}) interface{}{
  data, _ := ioutil.ReadFile(path)
  yaml.Unmarshal(data, t)
  return t
}