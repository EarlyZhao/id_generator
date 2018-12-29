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
  business_type := l.MustGetString("business_type", "business_type Must be valid")
  business_desc := l.Context.Input.GetString("business_desc", "")

  enable := l.MustGetString("enable", "start_at: need a number")
  list := &models.List{}
  models.DB.Where("business_type = ?", business_type).First(list)
  // disable
  if enable == "0"{
    list.Enable = false
  }else{
    interval := l.MustGetInt("interval", "interval: need a number")
    start_at := l.MustGetInt("start_at", "start_at: need a number")

    if list.ID == 0{
      l.Data["json"] = helpers.NewErrorRet(1003, "business_type Not Found")
      return
    }

    if list.StartedAt >= start_at || list.EndedAT >= start_at + interval{
      ret := fmt.Sprintf("start_at must greater than the old(%d), or the interval too small(old: %d)", list.StartedAt, list.Interval)
      l.Data["json"] = helpers.NewErrorRet(1004, ret)
      return
    }

    list.BusinessType = business_type
    list.BusinessDesc = business_desc
    list.Interval = interval
    list.StartedAt = start_at
    list.EndedAT = start_at + interval
    list.Enable = true
  }

  if err := models.DB.Save(list).Error; err != nil{
    l.Data["json"] = helpers.NewErrorRet(1002, err.Error())
    return
  }

  l.Data["json"] = helpers.NewSuccessRet(list)
  unique_id.UpdateWareHouse(list.BusinessType)
}