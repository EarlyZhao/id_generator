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
  buffersCount := conf.BufferCount()
  fmt.Println("\nBuffer count:", buffersCount)
  for _, list := range(lists){
    set := GetNewIdSet(buffersCount, list)
    House.AddNewToWareHouse(list.BusinessId(), set)
  }
}

func UpdateWareHouse(business_type string){
  list := &models.List{}

  models.DB.Where("business_type = ?", business_type).First(list)
  UpdateIdSet(list)
}

func GetUniqueId(business string) (uint64, error) {
  // acquire an id from data set
  id, err := House.Acquire(business)
  return id, err
}

func UpdateIdSet(list BufferInterface){
  if list.Usable(){
    buffersCount := conf.BufferCount()
    if set, ok := House.Get(list.BusinessId());ok{
      set.Reload(buffersCount, list)
    }else{
      set = GetNewIdSet(buffersCount, list)
      House.AddNewToWareHouse(list.BusinessId(), set)
    }
  }else{
    House.RemoveToWareHouse(list.BusinessId())
  }
}