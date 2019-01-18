package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go-boilerplate/models"
	"go-boilerplate/user"
)

type AuthMiddleware struct {
	userUCase *user.UserUseCase
}

func (rm *AuthMiddleware) SetValidateUser (next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwtConfig := c.Get("jwtConfig").(*jwt.Token)
		claims := jwtConfig.Claims.(*models.JWTClaims)
		user := rm.userUCase.GetUserById(claims.Id)

		if user ==nil {
			return nil
		}
		c.Set("user", user)
		return next(c)
	}
}

func NewAuthMiddleware (uuc *user.UserUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		userUCase:uuc,
	}
}

