package controllers

import(
  "github.com/id_generator/models"
  "fmt"
)

type ListController struct{
  Controller
}

func (l *ListController) Create(){
  fmt.Println(l.Context.Input.Params)
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
  }

  models.DB.Save(list)
}