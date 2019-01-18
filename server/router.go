package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-boilerplate/config"
	"go-boilerplate/datastore"
	"go-boilerplate/models"
	"go-boilerplate/user"
	"net/http"
)

func InitRouter() {
	//Get initialized db object
	dbCon := datastore.GetDBConnection()

	//Init Repositories
	sqlUserRepo := user.GetNewSQLUserRepository(dbCon)

	//Init UseCases
	userUseCase := user.GetNewUserUseCase(sqlUserRepo)

	//Initialize handlers and middleware
	userHandler := user.NewUserHandler(userUseCase)
	authMiddleware := NewAuthMiddleware(userUseCase)

	//Serve static files from the public directory
	EchoCon.Static("/public", "public")

	//Public routes
	EchoCon.GET("/api/sayhello", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Hello World")
	})

	//Auth routes
	EchoCon.GET("/api/OAuth/auth", userHandler.LoginHandler)
	EchoCon.GET("/api/OAuth/auth/callback", userHandler.LoginCallbackHandler)

	/*
	Protected routes for user
	 */
	r := EchoCon.Group("/api/user")
	r.Use(middleware.JWTWithConfig(getJWTConfig()))
	r.Use(authMiddleware.SetValidateUser)
}

func getJWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &models.JWTClaims{},
		SigningKey: []byte(config.GetConfig("JWT_SECRET")),
		TokenLookup:"cookie:" + config.GetConfig("JWT_COOKIE_NAME"),
		ContextKey: "jwtConfig",
	}
}