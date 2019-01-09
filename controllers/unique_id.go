package controllers

import(
    "github.com/id_generator/services/unique_id"
    "github.com/id_generator/helpers"
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