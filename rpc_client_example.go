// This is an example for client who will call the `MakeUniqueId`

package client_example

import(
  grpc "google.golang.org/grpc"
  "github.com/id_generator/grpc/id_rpc"
  ctx "context"
  "fmt"
  "os"
)

func Example() {
  // Set up a connection to the server.
  conn, err := grpc.Dial("localhost:1314", grpc.WithInsecure())
  if err != nil {
    fmt.Printf("did not connect: %v", err)
  }
  defer conn.Close()
  c := id_rpc.NewUniqueIdServiceClient(conn)

  // Contact the server and print out its response.
  name := "test"
  if len(os.Args) > 1 {
      name = os.Args[1]
  }

  id, err := c.MakeUniqueId(ctx.Background(), &id_rpc.BusinessType{Name: name})
  if err != nil {
    fmt.Printf("could not greet: %v", err)
  }else{
    fmt.Println(id.Id)
  }

}