package models

import(
  "time"
  "github.com/jinzhu/gorm"
)

type List struct{
  ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id"`

  BusinessType string `gorm:"unique;not null;" json:"business_type"`
  BusinessDesc string `json:"business_desc"`

  Interval uint64 `gorm:"DEFAULT:10000;" sql:"type:bigint;" json:"interval"`
  StartedAt uint64  `gorm:"DEFAULT:1;" sql:"type:bigint;" json:"started_at"`
  EndedAT uint64  `sql:"type:bigint;" json:"ended_at"`

  UpdatedAt  time.Time
  CreatedAt  time.Time

  Enable bool `gorm:"DEFAULT:1;" json:"enable"`
}

func GetAllList() []*List{
  lists := make([]*List, 1, 5)
  DB.Where("Enable = ?", true).Find(&lists)

  // var types []string

  // for _, list := range(lists){
  //  types = append(types, list.BusinessType)
  // }

  return lists
}

func (l *List) Usable() bool{
  if DB.NewRecord(l){
    return false
  }

  DB.Where("id = ?", l.ID).First(l)
  return l.Enable == true
}

func (l *List) Duration() uint64{
  return l.Interval
}

func (l *List) StartAt() uint64{
  return l.StartedAt
}

func (l *List) EndAT() uint64{
  return l.EndedAT
}

func (l *List) Update(){
  tx := DB.Begin()
  DB.Model(l).UpdateColumn("EndedAT", gorm.Expr("EndedAT + ?", l.Interval))
  tx.Commit()
}
