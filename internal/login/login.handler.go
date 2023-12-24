package login

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isd-sgcu/oph66-backend/internal/register"
	"golang.org/x/oauth2"
)

type Handler interface {
	GoogleLogin(c *gin.Context)
	GoogleCallback(c *gin.Context)
}

type handlerImpl struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return &handlerImpl{svc}
}

func (h *handlerImpl) GoogleLogin(c *gin.Context) {
	url := register.ExportedOAuth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *handlerImpl) GoogleCallback(c *gin.Context) {
	code := c.Query("code")

	token, err := register.ExportedOAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
		return
	}

	userInfo, err := getUserInfoFromGoogle(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	email, _ := userInfo["email"].(string)
	c.JSON(http.StatusOK, gin.H{"username": email, "token": token})
}

func getUserInfoFromGoogle(accessToken string) (map[string]interface{}, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
