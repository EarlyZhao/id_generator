package controllers



type UniqueIdController struct{
  Controller
}

func (u *UniqueIdController) Create(){
  u.Data["json"] = u.Context.Input.Params
}

func (u *UniqueIdController) Update(){
  u.Data["json"] = u.Context.Input.Params
  u.Context.Output.SetHeader("method", "Update")
}