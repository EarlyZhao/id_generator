package dispatcher

import(
  "testing"
  "github.com/EarlyZhao/id_generator/test"
  "github.com/EarlyZhao/id_generator/context"
  )


type TestController struct{
  Tag string
}

func (c *TestController) Create(){}


// func (c *TestController) Init(context *context.Context){}
// func (c *TestController) Serve(){}
// func (c *TestController) Done(){}
// func (c *TestController) RecoverFunc(context *context.Context){}

func TestRouterAddQuery(t *testing.T){
  var err error
  var findController interface{}
  var handlerMethod string
  dispatcher := NewDispatcher()

  controllerTag1 := "testController001"
  controllerTag2 := "testController002"
  controllerTag3 := "testController003"
  controllerTag4 := "testController004"
  path1 := "/test_path"
  path2 := "/test_path/:id/hahaha"
  path3 := "/test_path/test"
  path4 := "/test_path1"
  controller1 := &TestController{Tag: controllerTag1}
  controller2 := &TestController{Tag: controllerTag2}
  controller3 := &TestController{Tag: controllerTag3}
  controller4 := &TestController{Tag: controllerTag4}

  err = dispatcher.AddRoute("GET", path1, "Create", controller1)
  test.MustEqual(t, err, nil, "AddRoute() failed", "")

  err = dispatcher.AddRoute("GET", path2, "Create", controller2)
  test.MustEqual(t, err, nil, "AddRoute() failed", "")

  err = dispatcher.AddRoute("GET", path3, "Create", controller3)
  test.MustEqual(t, err, nil, "AddRoute() failed", "")

  err = dispatcher.AddRoute("GET", path4, "Create", controller4)
  test.MustEqual(t, err, nil, "AddRoute() failed", "")

  findController, _,handlerMethod, err = dispatcher.routes.FindHandler("GET" + path1)
  test.MustEqual(t, err, nil, "FindHandler() failed","")
  test.MustEqual(t, findController.(*TestController).Tag, controllerTag1, "FindHandler() failed","")
  test.MustEqual(t, handlerMethod, "Create", "FindHandler() failed when get right MethodName","")

  params := make(map[string]string)
  path2Test := "/test_path/lalala/hahaha" // path2
  findController, params,_, err = dispatcher.routes.FindHandler("GET" + path2Test)
  test.MustEqual(t, err, nil, "FindHandler() failed","")
  test.MustEqual(t, findController.(*TestController).Tag, controllerTag2, "FindHandler() failed","")
  test.MustEqual(t, params["id"], "lalala", "FindHandler() failed when get params from url","")

  findController, _,_, err = dispatcher.routes.FindHandler("GET" + path3)
  test.MustEqual(t, err, nil, "FindHandler() failed","")
  test.MustEqual(t, findController.(*TestController).Tag, controllerTag3, "FindHandler() failed","")

  findController, _,_, err = dispatcher.routes.FindHandler("GET" + path4)
  test.MustEqual(t, err, nil, "FindHandler() failed","")
  test.MustEqual(t, findController.(*TestController).Tag, controllerTag4, "FindHandler() failed","")


  err = dispatcher.AddRoute("POST", path1, "Create", controller1)
  test.MustEqual(t, err, nil, "AddRoute() failed", "")
  findController, _,_, err = dispatcher.routes.FindHandler("POST" + path1)
  test.MustEqual(t, findController.(*TestController).Tag, controllerTag1, "FindHandler() failed","")

  err = dispatcher.AddRoute("PUT", path2, "Create", controller2)
  test.MustEqual(t, err, nil, "AddRoute() failed", "")
  findController, _,_, err = dispatcher.routes.FindHandler("PUT" + path2)
  test.MustEqual(t, findController.(*TestController).Tag, controllerTag2, "FindHandler() failed","")
}

func TestNoMethod(t *testing.T){
  var hasError bool
  path1 := "/test_path1"
  dispatcher := NewDispatcher()
  defer func(){
    if e := recover(); e != nil{
      hasError = true
    }

    test.MustEqual(t, hasError, true, "AddRoute() could register method that not exist", "")
  }()

  dispatcher.AddRoute("GET", path1, "aMethodNotExist", &TestController{})
}

func TestPathConflict(t *testing.T){
  var hasError bool
  path1 := "/test_path1/test"
  path2 := "/test_path1/:id"
  dispatcher := NewDispatcher()

  dispatcher.AddRoute("GET", path1, "Create", &TestController{})
  defer func(){
    if e := recover(); e != nil{
      hasError = true
    }

    test.MustEqual(t, hasError, true, "AddRoute() could register conflict path", "")
  }()

  dispatcher.AddRoute("GET", path2, "Create", &TestController{})
}