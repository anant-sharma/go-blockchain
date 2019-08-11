package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"

	Config "github.com/anant-sharma/go-blockchain-config"
	jwt "github.com/dgrijalva/jwt-go"
)

/*
	Get Application Configuration
*/
var config = Config.GetConfig()

var myJwtSigningKey = []byte(config.Jwt.Secret)

/*
GenToken - A Util function to generate jwtToken which can be used in the request header
*/
func GenToken(id uint) (string, error) {

	/* Create Token */
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	/* Set token claims */
	jwtToken.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Second * config.Jwt.ExpiresIn).Unix(),
		"iss": config.Jwt.Issuer,
	}

	/* Sign the token with our secret */
	return jwtToken.SignedString(myJwtSigningKey)

}

// Sha256 to generate sha256
func Sha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateUUID - Function to generate UUID
func GenerateUUID() string {
	return uuid.New().String()
}
