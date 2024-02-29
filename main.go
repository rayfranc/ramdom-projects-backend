package main

import (
	"main/config"
	"main/controller"
	"main/helper"
	"main/middlewares"
	model "main/models"
	"main/repository"
	"main/routes"
	"main/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main(){
 

  db:=config.DatabaseConnection()
  validate:=validator.New()
  db.Table("people").AutoMigrate(&model.Person{})
  app:=gin.Default()
  app.Use(gin.CustomRecovery(middlewares.ErrorHandler))
	personRepository:=repository.NewPersonRepositoryImpl(db)
	personService:=services.NewPersonServiceImpl(personRepository,validate)
	personController:=controller.NewPersonController(personService)

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