package unique_id

import(
  "github.com/id_generator/models"
)

// generate ID uniquely and increasingly
// the ID is an integer that created from database
var House *WareHouse

func init(){
  go initWareHouse()
}

func initWareHouse(){
  <- models.ConnectionSucess
  // get business types
  lists := models.GetAllList()
  // create id set for every business
  House = NewWareHouse()
  for _, list := range(lists){
    business_type := list.BusinessType
    count := 2 // todo: from config file
    set := GetNewIdSet(count, list)
    House.SetHouse(business_type, set)
  }
}

func GetUniqueId(business string) (uint64, error) {
  // acquire a id from data set
  id, err := House.Acquire(business)
  return id, err
}