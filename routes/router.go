package routes


import(
        "id_generator/app"
        "id_generator/controllers"
)



func init(){
  app.AddRoute("GET", "/unique_ids/:id", "Create",&controllers.UniqueIdController{})
  app.AddRoute("GET", "/unique_ids/:id/update", "Update",&controllers.UniqueIdController{})
  app.AddRoute("GET", "/unique_ids/:id/update/:level", "Update",&controllers.UniqueIdController{})
}