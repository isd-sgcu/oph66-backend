package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"github.com/isd-sgcu/oph66-backend/internal/auth"
	"github.com/isd-sgcu/oph66-backend/utils"
	"google.golang.org/api/idtoken"
)

type AuthMiddleware gin.HandlerFunc

func NewAuthMiddleware(userRepo auth.Repository, cfg *cfgldr.Config) AuthMiddleware {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.ReturnError(c, apperror.InvalidToken)
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := idtoken.Validate(c, tokenString, cfg.OAuth2Config.ClientId)

		if err != nil {
			SecretKey := cfg.JWTConfig.SecretKey
			staffToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(SecretKey), nil
			})
			if err != nil {
				utils.ReturnError(c, apperror.InvalidToken)
				c.Abort()
				return
			}
			if staffToken.Valid && staffToken.Claims.(jwt.MapClaims)["role"] == "staff" {
				c.Set("faculty", staffToken.Claims.(jwt.MapClaims)["faculty"])
				c.Set("department", staffToken.Claims.(jwt.MapClaims)["department"])
				c.Set("faculty-wide", staffToken.Claims.(jwt.MapClaims)["faculty-wide"])
				c.Next()
				return
			} else {
				utils.ReturnError(c, apperror.InvalidToken)
				c.Abort()
				return
			}
		}

		if email, ok := token.Claims["email"].(string); ok {
			c.Set("email", email)
			c.Next()
			return
		} else {
			utils.ReturnError(c, apperror.ServiceUnavailable)
			c.Abort()
			return
		}
	}
}
