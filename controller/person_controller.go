package controller

import (
	"main/data/exceptions"
	"main/data/request"
	"main/data/response"
	"main/helper"
	"main/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonController struct{
	personService services.PersonService
}

func NewPersonController(ps services.PersonService) *PersonController {
	return &PersonController{personService: ps}
}

func (controller *PersonController)Create(ctx *gin.Context){
  createPersonRequest:=request.CreatePersonRequest{}
   err:=  ctx.ShouldBindJSON(&createPersonRequest);
   exceptions.ThrowBadRequest(ctx,err)
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