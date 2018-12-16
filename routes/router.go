package routes


import(
        "github.com/id_generator/app"
        "github.com/id_generator/controllers"
)



func init(){
  app.AddRoute("GET", "/unique_ids/:id", "Create",&controllers.UniqueIdController{})
  app.AddRoute("GET", "/unique_ids/:id/update", "Update",&controllers.UniqueIdController{})
  app.AddRoute("GET", "/unique_ids/:id/update/:level", "Update",&controllers.UniqueIdController{})
}