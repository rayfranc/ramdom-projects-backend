package controller

import (
	"fmt"
	"main/data/request"
	"main/data/response"
	"main/services"
	helper "main/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PersonController struct{
	personService services.PersonService
	validate         *validator.Validate
}

func NewPersonController(ps services.PersonService, validate *validator.Validate) *PersonController {
	return &PersonController{personService: ps,validate: validate,}
}

func (controller *PersonController)Create(ctx *gin.Context){
  createPersonRequest:=request.CreatePersonRequest{}
   err:=  ctx.ShouldBindJSON(&createPersonRequest);
   
   if err!=nil{
	
	ctx.AbortWithStatusJSON(http.StatusBadRequest,err.Error())	
	return
   }
    fmt.Println(createPersonRequest)
	fmt.Printf("Incoming Data %+v ", createPersonRequest)
  controller.personService.Create(createPersonRequest)
  SendJSONResponse(ctx,nil)
}

func (controller *PersonController)Update(ctx *gin.Context){
  updatePersonRequest:=request.UpdatePersonRequest{}
	err:=ctx.ShouldBindJSON(&updatePersonRequest)
	helper.ErrorPanic(err)
	personId:=ctx.Param("personId")
	id,err:= strconv.Atoi(personId)
	helper.ErrorPanic(err)
	updatePersonRequest.Id=id
	controller.personService.Update(updatePersonRequest)
	SendJSONResponse(ctx,nil)

}

func (controller *PersonController)FindAll(ctx *gin.Context){
	personsResponse:=controller.personService.FindAll()
	SendJSONResponse(ctx,personsResponse)
}

func (controller *PersonController)FindById(ctx *gin.Context){
	personId:=ctx.Param("personId")
	id, err :=strconv.Atoi(personId)
	 helper.ErrorPanic(err)
	 personResponse:=controller.personService.FindById(id)
	 SendJSONResponse(ctx,personResponse)
}

func (controller *PersonController)Delete(ctx *gin.Context){
     personId:=ctx.Param("personId")
	 id, err :=strconv.Atoi(personId)
	 helper.ErrorPanic(err)
	 controller.personService.Delete(id)

	 SendJSONResponse(ctx,nil)
}

func (controller *PersonController)Shuffle(ctx *gin.Context){
	shuffleRequest:=request.ShuffleRequest{}
	err:=ctx.ShouldBindJSON(&shuffleRequest)
	if err!=nil{
		fmt.Println(err)
	ctx.AbortWithStatusJSON(http.StatusBadRequest,err.Error())
	return
	}
	res:=controller.personService.Shuffle(shuffleRequest)
	ctx.Header("Content-Type","application/json")
	ctx.JSON(http.StatusOK,res)
}


func SendJSONResponse(ctx *gin.Context, data interface{}){
	if data!=nil{
	webResponse:=response.Response{
		Code: 200,
		Status:  "OK",
		Data:   data,
	}
	ctx.Header("Content-Type","application/json")
	ctx.JSON(http.StatusOK,webResponse)
	return
   }else{
	webResponse:=response.Response{
		Code: 200,
		Status:  "OK",
	}
	ctx.Header("Content-Type","application/json")
	ctx.JSON(http.StatusOK,webResponse)
   }
}