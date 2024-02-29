package exceptions

import (
	"main/data/response"
	"net/http"

	"github.com/gin-gonic/gin"
)



func ThrowBadRequest(ctx *gin.Context){
	
	webResponse:=response.Response{
		Code: 400,
		Status:  "Bad Request",
		Data: "Validation Error",
	}
	ctx.Header("Content-Type","application/json")
	ctx.JSON(http.StatusBadRequest,webResponse)
}