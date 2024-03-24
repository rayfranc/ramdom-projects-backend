package main

import (
	"main/controller"
	"main/routes"
	"main/services"
	helper "main/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main(){
 

//   db:=config.DatabaseConnection()
validate:=validator.New()
//   db.Table("people").AutoMigrate(&model.Person{})
  app:=gin.Default()
  //app.Use(gin.CustomRecovery(middlewares.ErrorHandler))
	// personRepository:=repository.NewPersonRepositoryImpl(db)
	 personService:=services.NewPersonServiceImpl()
	personController:=controller.NewPersonController(personService,validate)
	app.POST("/api/shuffle",personController.Shuffle)
app = routes.PersonRoutes(personController,app)
	
  
 server := &http.Server{
	Addr:           ":8080",
	Handler:        app,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
 }
  
  err := server.ListenAndServe() 
  helper.ErrorPanic(err)
  
}