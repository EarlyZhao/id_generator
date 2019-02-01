package unique_id

import(
  "github.com/EarlyZhao/id_generator/models"
  "github.com/EarlyZhao/id_generator/conf"
  "fmt"
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
  buffersCount := conf.Config.Buffers
  fmt.Println("\nBuffer count:", buffersCount)
  for _, list := range(lists){
    set := GetNewIdSet(buffersCount, list)
    House.AddNewToWareHouse(list.BusinessType, set)
  }
}

func UpdateWareHouse(business_type string){
  list := &models.List{}

  models.DB.Where("business_type = ?", business_type).First(list)
  if list.Usable(){
    if set, ok := House.HouseMap[business_type];ok{
      set.Reload(2, list)
    }else{
      set = GetNewIdSet(2, list)
      House.AddNewToWareHouse(list.BusinessType, set)
    }
  }else{
    House.RemoveToWareHouse(list.BusinessType)
  }
}

func GetUniqueId(business string) (uint64, error) {
  // acquire an id from data set
  id, err := House.Acquire(business)
  return id, err
}