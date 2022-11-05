package handler

import (
	"fmt"
	"net/http"

	"github.com/WorkWorkWork-Team/gov-voter-api/config"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func AuthorizeJWT(jwtService service.JWTService, appConfig config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < len(BEARER_SCHEMA) {
			errMessage := "Header is not containing any Bearer key."
			logrus.Warn(errMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errMessage,
			})
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			logrus.Error(err, ", Token: ", tokenString)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		if !token.Valid {
			errMessage := "Cannot validate token"
			logrus.Warn(errMessage, ", Token: ", tokenString)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errMessage,
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if claims["iss"] != appConfig.JWT_ISSUER {
			errMessage := "Token signed from unknown issuer"
			logrus.Warn(errMessage, ", Token: ", tokenString)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errMessage,
			})
		}
		logrus.Debug(claims)
		c.AddParam("CitizenID", fmt.Sprint(claims["CitizenID"]))
	}
}
