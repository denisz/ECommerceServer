package cart

import (
	"net/http"
	"store/controllers/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"store/controllers/catalog"
)

var SecretKey = []byte("sfvDUPC0Cj")
const CookieName = "CART_ID"

type Controller struct {
	common.Controller
}

func ReadFromRequest(c *gin.Context) *Session {
	session := &Session{}
	tokenString, err := c.Cookie(CookieName)
	if err != nil {
		return session
	}

	jwt.ParseWithClaims(tokenString, session, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	return session
}

func WriteToResponse(c *gin.Context, session *Session) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), session)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.SetCookie(CookieName, tokenString, 7 * 24 * 3600, "/","",false, true)
}

func (p *Controller) IndexPOST(c *gin.Context) {
	session := ReadFromRequest(c)
	c.JSON(http.StatusOK, &Cart{
		Positions: session.Positions,
	})
}

func (p *Controller) DetailPOST(c *gin.Context) {
	session := ReadFromRequest(c)
	cart := &Cart{}

	for _, v := range session.Positions {
		var product catalog.Product
		err := p.GetStoreNode().One("ID", v.ProductID, &product)
		if err != nil { continue }
		cart.Products = append(cart.Products, product)
		cart.Positions = append(cart.Positions, v)
	}

	c.JSON(http.StatusOK, cart)
}

/// Обновить позицию
func (p *Controller) UpdatePOST(c *gin.Context) {
	var json ItemDTO

	if err := c.ShouldBindJSON(&json); err == nil {
		session := ReadFromRequest(c)
		var positions []Position

		for _, v := range session.Positions {
			if v.ProductID == json.ProductID {
				v.Amount = json.Amount
			}

			if v.Amount > 0 {
				positions = append(positions, v)
			}
		}

		session.Positions = positions
		WriteToResponse(c, session)
		c.JSON(http.StatusOK, &Cart{
			Positions: positions,
		})
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

//Добавить позицию
func (p *Controller) InsertPOST(c *gin.Context) {
	var json ItemDTO

	if err := c.ShouldBindJSON(&json); err == nil {
		exists := false

		session := ReadFromRequest(c)
		var positions []Position

		for _, v := range session.Positions {
			if v.ProductID == json.ProductID {
				v.Amount = v.Amount + json.Amount
				exists = true
			}
			positions = append(positions, v)
		}

		if !exists {
			positions = append(positions, Position{
				ProductID: json.ProductID,
				Amount: json.Amount,
			})
		}

		session.Positions = positions
		WriteToResponse(c, session)
		c.JSON(http.StatusOK, &Cart{
			Positions: positions,
		})
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}


