package routes


import(
        "github.com/id_generator/app"
        "github.com/id_generator/controllers"
)



func init(){
  app.AddRoute("GET", "/unique_ids/:id", "Create",&controllers.UniqueIdController{})
  app.AddRoute("GET", "/lists", "Index", &controllers.ListController{})
  app.AddRoute("POST", "/lists", "Create", &controllers.ListController{})
  app.AddRoute("PUT", "/lists", "Update", &controllers.ListController{})
}