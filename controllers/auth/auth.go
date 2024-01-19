package AuthController

import (
	// "encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"web-service/dto/auth"
	ResponseError "web-service/utils/errors"
)

func Login(c *gin.Context) {
	var input AuthDTO.LoginRequestDTO

	if err := c.ShouldBindJSON(&input); err != nil {

		errors := ResponseError.ParseError(err)

		// customErrors := ResponseError.MapValidationErrors(err)
		// customErrorsJson, _ := json.Marshal(customErrors)
		log.Error("Error: ", errors)

		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	c.JSON(200, gin.H{"message": "Login success"})
}
