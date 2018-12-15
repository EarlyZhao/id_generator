package dispatcher

import(
  "strings"
  "errors"
)

type Node struct{
  Children map[string]*Node
  Name string

  Parameter bool // if the parameter node, like /:id
  ParameterName string
  HasParameterNode bool
  // ParameterValue string

  PathEnd bool // the endpoint of path, ex: the C is the endpoint of path /a/b/C
  Value interface{} // controller that deal the bussiness logic
  HandleMethod string
  // T interface{} // type of Value
}

func CreateNode(key_value *string) *Node{
  new_node := &Node{
        Name: *key_value,
        Children: make(map[string]*Node),
        PathEnd: false,
        Parameter: false,
  }

  if strings.Contains(*key_value, ":"){
    new_node.Parameter = true
    new_node.ParameterName = strings.Trim(*key_value, ":")
  }

  return new_node
}

func (n *Node) ParameterNode() (*Node, error){
  var node *Node
  if n.HasParameterNode == false{
    return node, errors.New("no parameter node")
  }

  for _, node = range(n.Children){
    if node.Parameter{
      return node, nil
    }
  }

  return node, errors.New("no parameter node")
}