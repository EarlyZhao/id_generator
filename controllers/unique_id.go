package controllers

import(
    "github.com/id_generator/services/unique_id"
)

type UniqueIdController struct{
  Controller
}

func (u *UniqueIdController) Create(){
  // u.Data["json"] = u.Context.Input.Params
  // var business string
  // business := u.Context.Input.Params["id"][0]
  business := u.Context.Input.GetString("id", "")
  id, err := unique_id.GetUniqueId(business)
  ret := make(map[string]interface{})

  if err == nil{
    ret["id"] = id
    ret["status"] = 0
  }else{
    ret["msg"] = err.Error()
    ret["status"] = 1
  }

  ret["business"] = business

  u.Data["json"] = ret
}

func (u *UniqueIdController) Update(){
  u.Data["json"] = u.Context.Input.Params
  u.Context.Output.SetHeader("method", "Update")
}