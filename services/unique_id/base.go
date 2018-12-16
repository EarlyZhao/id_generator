package unique_id

var House *WareHouse

func init(){
  // get business types
  types := []string{"test"}
  // create id set for every business
  House = NewWareHouse()
  for _, business := range(types){
    set := GetNewIdSet()
    House.SetHouse(business, set)
  }

}

func GetUniqueId(business string) (int, error) {
  // acquire a id from data set
  id, err := House.Acquire(business)
  return id, err
}