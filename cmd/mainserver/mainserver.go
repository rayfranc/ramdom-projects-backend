package main

import (
	"main/controller"
	config "main/db"
	model "main/models"
	"main/repository"
	"main/routes"
	"main/services"
	helper "main/utils"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main(){
 

  db:=config.DatabaseConnection()
  validate:=validator.New()
  db.Table("people").AutoMigrate(&model.Person{})
  app:=gin.Default()
  config := cors.DefaultConfig()
  config.AllowAllOrigins = true
  app.Use(cors.New(config))
	personRepository:=repository.NewPersonRepositoryImpl(db)
	personService:=services.NewPersonServiceImpl(personRepository)
	personController:=controller.NewPersonController(personService,validate)
	app.POST("/api/shuffle",personController.Shuffle)
app = routes.PersonRoutes(personController,app)
	
  
 server := &http.Server{
	Addr:           ":8888",
	Handler:        app,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
 }
  
  err := server.ListenAndServe() 
  helper.ErrorPanic(err)
  
}