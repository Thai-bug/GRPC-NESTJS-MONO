package AuthController

import (
	"encoding/json"
	"net/http"
	"web-service/dto/auth"
	ResponseError "web-service/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/pspiagicw/colorlog"
)

func Login(c *gin.Context) {
	var input AuthDTO.LoginRequestDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		customErrors := ResponseError.MapValidationErrors(err)
		customErrorsJson, _ := json.Marshal(customErrors)
		colorlog.LogError("Error: ", string(customErrorsJson))

		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors})
		return
	}

	c.JSON(200, gin.H{"message": "Login success"})
}
