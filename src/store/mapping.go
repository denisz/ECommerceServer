package store

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	Middleware "store/middlewares"
	"store/services/router"
	"store/services/api"
)

var I = Middleware.Instrument

func CreateMapping(api *api.API, allowOrigins []string) http.Handler {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Handlers
	h := router.NewRouter(api)

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
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
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
			//TODO:Добавить нормальную проверку пользователя
			if userId == "test" && password == "test" {
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

	//Mapping
	v1 := r.Group("/api/v1/")
	{
		v1.POST("/settings", I("/settings"), h.SettingsIndex)
		v1.POST("/catalog/collections", I("/catalog/collection"), h.CatalogCollectionsPOST)
		v1.POST("/catalog/collection/:sku", I("/catalog/collection/:sku"), h.CatalogCollectionDetailPOST)
		v1.POST("/catalog/products/:sku", I("/catalog/products/:sku"), h.CatalogProductsPOST)
		v1.POST("/catalog/product/:sku", I("/catalog/product/:sku"), h.CatalogProductDetailPOST)
		v1.POST("/catalog/notation/:sku", I("/catalog/notation/:sku"), h.CatalogNotationPOST)
		v1.POST("/catalog/search", I("/catalog/search"), h.CatalogSearchProductsPOST)
		v1.POST("/sales", I("/sales"), h.SalesIndexPOST)
		v1.POST("/cart", I("/cart"), h.CartIndexPOST)
		v1.POST("/cart/detail", I("/cart"), h.CartDetailPOST)
		v1.POST("/cart/update", I("/cart/update"), h.CartUpdatePOST)
		v1.POST("/cart/address", I("/cart/address"), h.CartUpdateAddressPOST)
		v1.POST("/cart/delivery", I("/cart/delivery"), h.CartUpdateDeliveryPOST)
		v1.POST("/cart/checkout", I("/order/checkout"), h.CartCheckoutPOST)
		v1.POST("/orders/check/:invoice", I("/orders/check/:invoice"), h.OrderDetailPOST)

		v1.POST("/account/login", authMiddleware.LoginHandler)
		v1.GET("/load/catalog", I("/load/catalog"), h.LoaderCatalogFromGoogle)
		v1.GET("/load/ads", I("/load/ads"), h.LoaderAdsFromGoogle)
	}

	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.POST("/account/me", I("/account/me"), h.AccountMePOST)
		v1.POST("/order/:id", I("/order/:id"), h.OrderUpdatePOST)
		v1.POST("/forms/order/:id", I("/forms/order/:id"), h.FormsOrderPOST)
		v1.POST("/forms/batch/:id", I("/forms/order/:id"), h.FormsBatchPOST)
		v1.POST("/orders/list", I("/orders/list"), h.OrderListPOST)
		v1.POST("/orders/clear", I("/orders/clear"), h.OrderClearExpiredPOST)
		v1.POST("/orders/search", I("/orders/search"), h.SearchOrdersPOST)
		v1.POST("/orders/batch", I("/orders/batch"), h.CreateBatchPOST)
		v1.POST("/batches/search", I("/batches/search"), h.SearchBatchesPOST)
		v1.GET("/batches/:id/checkin", I("/batches/:id/checkin"), h.CheckInBatchGET)
		v1.DELETE("/batch/:id", I("/batch/:id"), h.BreakBatchDELETE)
		v1.GET("/batch/:id", I("/batch/:id"), h.BatchDetailGET)
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	return r
}
