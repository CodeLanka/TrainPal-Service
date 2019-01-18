package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"go-boilerplate/config"
	"go-boilerplate/models"
	"golang.org/x/oauth2"
	"net/http"
	"strconv"
	"time"
)

var (
	googleOauthConfig *oauth2.Config
	/*
	Set some random string for each request
	*/
	oauthStateString = "random"
)

type UHandler struct {
	userUseCase *UserUseCase
}

func NewUserHandler(u *UserUseCase) *UHandler {
	googleOauthConfig = getGOAuthConfig()
	return &UHandler {
		userUseCase: u,
	}
}

func (userHandler* UHandler) LoginHandler(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (userHandler* UHandler) LoginCallbackHandler(c echo.Context) error {
	/*
	Get OAuth config from here
	*/
	user := new(models.User)
	tkn, e := getJWToken(user)

	if e != nil {
		log.Warn("Trouble converting json", e)
	}

	c.SetCookie(getCookie(tkn))

	return c.String(http.StatusOK, tkn)
}

func getJWToken(u *models.User) (string, error) {
	expTime, _ := strconv.Atoi(config.GetConfig("JWT_TOKEN_EXP_TIME_HOURS"))
	claims := &models.JWTClaims {
		Name: u.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expTime) * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetConfig("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getGOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
	}
}

func getCookie(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = config.GetConfig("JWT_COOKIE_NAME")
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Secure = true
	expTime, _ := strconv.Atoi(config.GetConfig("JWT_COOKIE_EXP_TIME_HOURS"))
	cookie.Expires = time.Now().Add(time.Duration(expTime) * time.Hour)
	return cookie
}