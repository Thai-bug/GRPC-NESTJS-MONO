package Middlewares

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/pspiagicw/colorlog"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := c.GetRawData()
		if err == nil {
			colorlog.LogInfo("Incoming Request: ",
				c.Request.Method,
				c.Request.URL,
				string(bodyBytes))
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()
	}
}
