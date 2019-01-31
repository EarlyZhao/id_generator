package dispatcher


import(
  "fmt"
  "strings"
  "errors"
  "reflect"
)
// TrieTree
type Tree struct{
  Root *Node
}



func (t *Tree) Add(path string, handleMethod string, handler interface{}) error{
  keys := CreateKeys(path)
  var pre *Node
  current := t.Root

  for _, key_value := range(keys){
    pre = current
    if child, ok := current.Children[key_value];ok{
      current = child
    }else{
      new_node := CreateNode(&key_value)
      current.Children[key_value] = new_node
      if new_node.Parameter{
        current.HasParameterNode = true
      }
      current = new_node
    }
  }
  // now , the current is the endpoint of path

  // check conflict, /a/b/c is conflict with /a/b/:id ,
  // the dispatcher cannot distinguish them
  for _, node := range(pre.Children){
    parameter_node := node.Parameter || current.Parameter
    if node.PathEnd && parameter_node{
      panic(fmt.Sprintf("There was a Path that conflict with %s \n", path))
    }
  }

  current.PathEnd = true
  current.Value = handler
  reflectValue := reflect.ValueOf(handler)
  reflectMethod := reflectValue.MethodByName(handleMethod)
  if reflectMethod.IsValid() == false{
    panic(fmt.Sprintf("%v Not Has Method: %s", handler, handleMethod))
  }
  current.HandleMethod = handleMethod
  return nil
}


func (t *Tree) FindHandler(path string) (interface{}, map[string]string, string, error){
  keys := CreateKeys(path)
  current := t.Root
  params := make(map[string]string)

  for _, key := range(keys){
    if child, ok := current.Children[key];ok{
      current = child
    }else if current.HasParameterNode{
      if child, err := current.ParameterNode(); err ==nil {
        params[child.ParameterName] = key
        current = child
      }else{
        return nil, params,"", errors.New("Not route1")
      }
    }else{
      return nil, params, "",errors.New("Not route2")
    }
  }

  if current.PathEnd == false{
    return nil, params, "", errors.New("Not route")
  }

  handler := current.Value
  return handler, params, current.HandleMethod, nil
}

func CreateKeys(path string) []string{
  path = strings.Trim(path, "/")
  keys := strings.Split(path, "/")
  return keys
}
