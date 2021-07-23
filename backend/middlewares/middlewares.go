package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var NoAuth = false

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !NoAuth {
			bearToken := c.Request.Header.Get("Authorization")
			strArr := strings.Split(bearToken, " ")
			if len(strArr) != 2 {
				c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
				c.Abort()
				return
			}

			token, err := jwt.Parse(strArr[1], func(token *jwt.Token) (interface{}, error) {
				//Make sure that the token method conform to "SigningMethodHMAC"
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("superDuperSaveRecipeDB"), nil
			})

			if err != nil {
				c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
				c.Abort()
				return
			}

			if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
				c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
