package controllers

import(
    "github.com/EarlyZhao/id_generator/services/unique_id"
    "github.com/EarlyZhao/id_generator/helpers"
)

type UniqueIdController struct{
  Controller
}

func (u *UniqueIdController) Create(){

  business := u.Context.Input.GetString("id", "")
  id, err := unique_id.GetUniqueId(business)
  ret := make(map[string]interface{})

  if err == nil{
    ret["business"] = business
    ret["id"] = id
    u.Data["json"] = helpers.NewSuccessRet(ret)
  }else{
    u.Data["json"] = helpers.NewErrorRet(1004, err.Error())
    return
  }

}

func (u *UniqueIdController) Update(){
  u.Data["json"] = u.Context.Input.Params
  u.Context.Output.SetHeader("method", "Update")
}