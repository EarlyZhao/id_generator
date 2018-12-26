
// 解耦 middleware 与 controller
package context



type Context struct{
  Input *Input
  Output *Output
}

func NewContext() *Context{
  return &Context{Input: &Input{}, Output: &Output{}}
}
