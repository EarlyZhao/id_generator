package models

import(
  "time"
  "github.com/jinzhu/gorm"
)

type List struct{
  ID int `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`

  BusinessType string `gorm:"unique;not null"`
  BusinessDesc string

  Interval uint64 `gorm:"DEFAULT:10000" sql:"type:bigint`
  StartedAt uint64  `gorm:"DEFAULT:1" sql:"type:bigint"`
  EndedAT uint64  `sql:"type:bigint"`

  UpdatedAt  time.Time

}

func GetAllList() []*List{
  lists := make([]*List, 1, 5)
  DB.Find(&lists)

  // var types []string

  // for _, list := range(lists){
  //  types = append(types, list.BusinessType)
  // }

  return lists
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
