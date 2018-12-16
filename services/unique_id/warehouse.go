package unique_id

import(
    "errors"
)

type WareHouse struct {
  HouseMap map[string]*IdSet // every business type has an IdSet
}

func (w *WareHouse) Acquire(business string) (uint64, error){
  businessSet, ok := w.HouseMap[business]
  if ok == false{
    return 0, errors.New("no business type")
  }

  id := businessSet.GetId()
  return id , nil
}

func (w *WareHouse) SetHouse(business string,  set*IdSet){
  // if w.HouseMap == nil{
  //  w.HouseMap = make(map[string]*IdSet)
  // }
  w.HouseMap[business] = set
}

func NewWareHouse() *WareHouse{
  return &WareHouse{HouseMap: make(map[string]*IdSet)}
}