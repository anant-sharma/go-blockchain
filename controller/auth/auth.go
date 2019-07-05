package authcontroller

import (
	"net/http"
	"time"

	utils "github.com/anant-sharma/go-blockchain/common"
	Config "github.com/anant-sharma/go-blockchain/config"
	"github.com/gin-gonic/gin"
)

/*
	Get Application Configuration
*/
var config = Config.GetConfig()

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
