package api

import (
	"net/http"
	//"gopkg.in/hlandau/passlib.v1"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	. "store/models"
)

type ControllerAccount struct {
	Controller
}


func (p *ControllerAccount) Index(w http.ResponseWriter, r *http.Request) {

}

func (p *ControllerAccount) LoginPOST(c *gin.Context) {
	var json LoginDTO
	if err := c.ShouldBindJSON(&json); err == nil {
		if ok, err := govalidator.ValidateStruct(json); !ok {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func (p *ControllerAccount) UpdatePOST(c *gin.Context) {

}

// Создаем нового пользователя и отправляем пароль на почту
func (p *ControllerAccount) RegisterPOST(c *gin.Context) {
	var json EmailRegisterDTO
	if err := c.ShouldBindJSON(&json); err == nil {
		if ok, err := govalidator.ValidateStruct(json); !ok {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

// Отправить пароль на почту пользователя
func (p *ControllerAccount) ResetPasswordPOST(c *gin.Context) {
	var json ForgetDTO
	if err := c.ShouldBindJSON(&json); err == nil {
		if ok, err := govalidator.ValidateStruct(json); !ok {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

