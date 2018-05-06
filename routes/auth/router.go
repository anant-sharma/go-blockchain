package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	utils "gitlab.com/anant-sharma/go-boilerplate/common"
	Config "gitlab.com/anant-sharma/go-boilerplate/config"
)

/*
	Get Application Configuration
*/
var config = Config.GetConfig()

/*
RegisterRouter Router Group
*/
func RegisterRouter(router *gin.RouterGroup) {
	router.POST("/", Authenticate)
}

/*
Authenticate function to authenticate user
*/
func Authenticate(c *gin.Context) {

	/* Generate Token */
	token, err := utils.GenToken(1)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":     token,
		"expiresAt": time.Now().Add(time.Second * config.Jwt.ExpiresIn).Unix(),
		"expiresIn": config.Jwt.ExpiresIn,
		"tokenType": "Bearer",
	})
}
