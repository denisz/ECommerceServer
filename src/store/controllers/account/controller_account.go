package account

import (
	"store/controllers/common"
	"net/http"
	//"gopkg.in/hlandau/passlib.v1"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	common.Controller
}


func (p *Controller) Index(w http.ResponseWriter, r *http.Request) {

}

func (p *Controller) LoginPOST(c *gin.Context) {
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

func (p *Controller) UpdatePOST(c *gin.Context) {

}

// Создаем нового пользователя и отправляем пароль на почту
func (p *Controller) RegisterPOST(c *gin.Context) {
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
func (p *Controller) ResetPasswordPOST(c *gin.Context) {
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

