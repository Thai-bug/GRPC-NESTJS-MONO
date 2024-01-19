package Middlewares

import (
	"bytes"
	// "encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := c.GetRawData()
		if err == nil {
			// body := json.RawMessage()
			log.Info("Incoming Request: ",
				c.Request.Method,
				c.Request.URL,
				string(bodyBytes))
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()
	}
}
