package store

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	Middleware "store/middlewares"
)

var I = Middleware.Instrument

func createRouter(store *Store) http.Handler {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Statistics middleware
	r.Use(Middleware.Stats())

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// CORS Request
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{store.Config.ServerURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "Russia zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if userId == "admin" && password == "admin" {
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header:Authorization",
		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",
		// TimeFunc provides the current time. You can override it to use another time value.
		// This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	v1 := r.Group("/api/v1/")
	{
		v1.POST("/settings", I("/settings"), store.Settings.Index)
		v1.POST("/catalog/collections", I("/catalog/collection"), store.Catalog.CollectionsPOST)
		v1.POST("/catalog/collection/:sku", I("/catalog/collection/:sku"), store.Catalog.CollectionPOST)
		v1.POST("/catalog/products/:sku", I("/catalog/products/:sku"), store.Catalog.ProductsPOST)
		v1.POST("/catalog/product/:sku", I("/catalog/product/:sku"), store.Catalog.ProductPOST)
		v1.POST("/sales", I("/sales"), store.Sales.IndexPOST)
		v1.POST("/cart", I("/cart"), store.Cart.IndexPOST)
		v1.POST("/cart/detail", I("/cart"), store.Cart.DetailPOST)
		v1.POST("/cart/update", I("/cart/update"), store.Cart.UpdatePOST)
		v1.POST("/cart/address", I("/cart/address"), store.Cart.UpdateAddressPOST)
		v1.POST("/cart/delivery", I("/cart/delivery"), store.Cart.UpdateDeliveryPOST)
		v1.POST("/order/checkout", I("/order/checkout"), store.Order.CheckoutPOST)

		v1.POST("/account/login", authMiddleware.LoginHandler)
		v1.GET("/load/catalog", I("/load/catalog"), store.Loader.CatalogFromGoogle)
		v1.GET("/load/ads", I("/load/ads"), store.Loader.AdsFromGoogle)
	}

	v1.Use(authMiddleware.MiddlewareFunc())
	{
		//v1.GET("/admin", )
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	return r
}