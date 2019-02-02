package controllers
// for smoking test

import(
  "github.com/EarlyZhao/id_generator/models"
  "github.com/EarlyZhao/id_generator/helpers"
)

type SmokingController struct{
  Controller
}

func (s *SmokingController)Smoking(){
  user := models.List{}
  err := models.DB.Take(&user).Error
  if err != nil{
    s.Data["json"] = helpers.NewErrorRet(1005, err.Error())
    s.Abort()
  }
}