package unique_id

import(
    "errors"
    "sync"
)

type WareHouse struct {
  HouseMap map[string]*IdSet // every business type has an IdSet
  mutex sync.Mutex
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
  w.mutex.Lock()
  defer w.mutex.Unlock()
  w.HouseMap[business] = set
}

func NewWareHouse() *WareHouse{
  return &WareHouse{HouseMap: make(map[string]*IdSet)}
}

func (w *WareHouse) RemoveToWareHouse(business_type string){
  w.mutex.Lock()
  defer w.mutex.Unlock()

  if _, ok := w.HouseMap[business_type];ok{
    delete(w.HouseMap, business_type)
  }
}

func (w *WareHouse) AddNewToWareHouse(business_type string, set *IdSet){
  House.SetHouse(business_type, set)
}

// func (w *WareHouse)