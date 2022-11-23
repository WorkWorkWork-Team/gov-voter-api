package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/WorkWorkWork-Team/gov-voter-api/config"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

var authBearerRegExp *regexp.Regexp = regexp.MustCompile("[B|b]earer (.*)")

func AuthorizeJWT(jwtService service.JWTService, appConfig config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStringList := authBearerRegExp.FindStringSubmatch(c.GetHeader("Authorization"))
		logrus.Debug("RegExp Debug: ", tokenStringList)
		if len(tokenStringList) != 2 {
			logrus.Warn("Authorization Header is malformed: ", c.GetHeader("Authorization"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token is malformed and request is unauthorized"})
			return
		}

		tokenString := tokenStringList[1]
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
