package middlewares

import (
	"main/data/response"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)



func ErrorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	switch goErr.TypeName(){
	case "validator.ValidationErrors":
		httpResponse := response.Response{Code: 400, Status: "Bad Request", Data: goErr.Error()}
		c.AbortWithStatusJSON(400, httpResponse)
	case "*strconv.NumError":
		httpResponse := response.Response{Code: 400, Status: "Bad Request", Data: goErr.Error()}
		c.AbortWithStatusJSON(400, httpResponse)
	default:
		httpResponse := response.Response{Code: 500, Status: goErr.TypeName(), Data: goErr.Error()}
		c.AbortWithStatusJSON(500, httpResponse)
	}
	
}