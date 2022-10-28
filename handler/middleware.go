package handler

import (
	"net/http"

	"github.com/WorkWorkWork-Team/common-go/jwtservice"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func AuthorizeJWT(jwtService jwtservice.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < len(BEARER_SCHEMA) {
			logrus.Warn("Header is not containing any Bearer key.")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			logrus.Warn("Cannot validate token ", tokenString)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		claims := token.Claims.(jwt.MapClaims)
		logrus.Info(claims)
	}
}
