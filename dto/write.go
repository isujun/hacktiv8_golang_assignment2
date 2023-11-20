package dto

import (
	"assignment2/dto/request_response"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, response *request_response.Response) {
	ctx.JSON(response.Status, response)
}