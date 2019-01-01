package routes


import(
        "github.com/id_generator/app"
        "github.com/id_generator/controllers"
)



func init(){
  app.AddRoute("GET", "/unique_ids/:id", "Create",&controllers.UniqueIdController{})
  app.AddRoute("GET", "/unique_ids/:id/update", "Update", &controllers.UniqueIdController{})
  app.AddRoute("POST", "/unique_ids/lists", "Create", &controllers.ListController{})
  app.AddRoute("PUT", "/unique_ids/lists", "Update", &controllers.ListController{})
}