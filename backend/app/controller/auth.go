package controller

import (
	"time"
	"yarujun/app/model"
	"yarujun/app/types"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func NewJwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour * 24,
		MaxRefresh: time.Hour * 24 * 7,
		SendCookie: false,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{
				jwt.IdentityKey: data,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var l types.LoginRequest
			if err := c.ShouldBind(&l); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			isValid, user_id := isValid(l)
			if !isValid {
				return "", jwt.ErrFailedAuthentication
			}

			return user_id, nil
		},
	})

	if err != nil {
		return nil, err
	}

	err = jwtMiddleware.MiddlewareInit()

	if err != nil {
		return nil, err
	}

	return jwtMiddleware, nil
}

func isValid(l types.LoginRequest) (bool, string) {
	password, user_id, err := model.GetLoginInfo(l.Name)
	if err != nil {
		return false, ""
	}

	return password == l.Password, user_id
}
