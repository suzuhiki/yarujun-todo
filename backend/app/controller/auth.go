package controller

import (
	"time"
	"yarujun/app/model"

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
			var l loginRequest

			if err := c.ShouldBind(&l); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if !l.isValid() {
				return "", jwt.ErrFailedAuthentication
			}

			return l.Name, nil
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

func (l loginRequest) isValid() bool {
	password := model.GetPassword(l.Name)

	return password == l.Password
}
