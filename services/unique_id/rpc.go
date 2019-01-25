package unique_id

import(
  "github.com/id_generator/grpc/id_rpc"
  "github.com/id_generator/logs"
  ctx "context"
  "fmt"
)


// type UniqueIdServiceServer interface {
//   GetUniqueId(context.Context, *BusinessType) (*UniqueId, error)
// }

type UniqueIdRpcService struct{}

func (u *UniqueIdRpcService) MakeUniqueId(c ctx.Context, business *id_rpc.BusinessType) (*id_rpc.UniqueId, error){
  unique_id := &id_rpc.UniqueId{}

  id, err := GetUniqueId(business.Name)
  if err != nil{
    logs.Info(fmt.Sprintf("(%s)MakeUniqueId error: %v", business.Name, err))
    return unique_id, err
  }

  unique_id.Id = id
  unique_id.BusinessType = business.Name

  return unique_id, nil
}

