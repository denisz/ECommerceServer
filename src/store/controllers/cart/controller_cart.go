package cart

import (
	"net/http"
	"store/controllers/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SecretKey = []byte("sfvDUPC0Cj")
const CookieName = "CART_ID"

type Controller struct {
	common.Controller
}

func ReadFromRequest(c *gin.Context) *Cart {
	cart := &Cart{}
	tokenString, err := c.Cookie(CookieName)
	if err != nil {
		return cart
	}

	jwt.ParseWithClaims(tokenString, cart, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	return cart
}

func WriteToResponse(c *gin.Context, cart *Cart) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), cart)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.SetCookie(CookieName, tokenString, 0, "/","",false, true)
}

func (p *Controller) DetailGET(c *gin.Context) {
	cart := ReadFromRequest(c)
	c.JSON(http.StatusOK, cart)
}

func (p *Controller) UpdatePOST(c *gin.Context) {
	var json CartDTO

	if err := c.ShouldBindJSON(&json); err == nil {
		cart := &Cart{ Items: []Item{}, }

		for _, v := range json.Items {
			cart.Items = append(cart.Items, Item{
				ProductID: v.ProductID,
				Amount: v.Amount,
			})
		}

		WriteToResponse(c, cart)
		c.JSON(http.StatusOK, cart)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func (p *Controller) InsertPOST(c *gin.Context) {
	var json ItemDTO

	if err := c.ShouldBindJSON(&json); err == nil {
		exists := false

		cart := ReadFromRequest(c)
		var items []Item

		for _, v := range cart.Items {
			if v.ProductID == json.ProductID {
				v.Amount = v.Amount + json.Amount
				exists = true
			}
			items = append(items, v)
		}

		if !exists {
			items = append(items, Item{
				ProductID: json.ProductID,
				Amount: json.Amount,
			})
		}

		cart.Items = items

		WriteToResponse(c, cart)
		c.JSON(http.StatusOK, cart)
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}


