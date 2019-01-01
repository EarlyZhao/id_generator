package controllers

import(
  "github.com/id_generator/models"
  "fmt"
  "github.com/id_generator/helpers"
  "github.com/id_generator/services/unique_id"
)

type ListController struct{
  Controller
}

func (l *ListController) Create(){
  business_type := l.MustGetString("business_type", "business_type Must be valid")
  business_desc := l.Context.Input.GetString("business_desc", "")
  interval := l.MustGetInt("interval", "interval: need a number")
  start_at := l.MustGetInt("start_at", "start_at: need a number")

  list := &models.List{
    BusinessType: business_type,
    BusinessDesc: business_desc,
    Interval: interval,
    StartedAt: start_at,
    EndedAT: start_at + interval,
    Enable: true,
  }

  if err := models.DB.Save(list).Error; err != nil{
    l.Data["json"] = helpers.NewErrorRet(1002, err.Error())
    return
  }

  l.Data["json"] = helpers.NewSuccessRet(list)
  unique_id.UpdateWareHouse(list.BusinessType)
}

func (l *ListController) Update(){
  var ret string
  business_type := l.MustGetString("business_type", "business_type Must be valid")
  business_desc := l.Context.Input.GetString("business_desc", "")

  enable := l.MustGetString("enable", "start_at: need a number")
  list := &models.List{}
  models.DB.Where("business_type = ?", business_type).First(list)
  // disable
  if enable == "0"{
    list.Enable = false
  }else{

    if list.ID <= 0{
      l.Data["json"] = helpers.NewErrorRet(1003, "business_type Not Found")
      return
    }

    if list.Enable == true{
      ret = "You must disable the bisiness type first"
      l.Data["json"] = helpers.NewErrorRet(1005, ret)
      return
    }

    ended_at := l.Context.Input.GetUint("ended_at", 0)
    if ended_at > 0 {
      if list.EndedAT >= ended_at {
        ret = fmt.Sprintf("ended_at must greater than the old(%d)", list.EndedAT)
        l.Data["json"] = helpers.NewErrorRet(1004, ret)
        return
      }else{
        list.EndedAT = ended_at
      }
    }

    interval := l.Context.Input.GetUint("interval", 0)
    if interval > 0{
      list.Interval = interval
    }

    if business_desc == ""{
      list.BusinessDesc = business_desc
    }

    list.Enable = true
  }

  if err := models.DB.Save(list).Error; err != nil{
    l.Data["json"] = helpers.NewErrorRet(1002, err.Error())
    return
  }

  l.Data["json"] = helpers.NewSuccessRet(list)
  unique_id.UpdateWareHouse(list.BusinessType)
}