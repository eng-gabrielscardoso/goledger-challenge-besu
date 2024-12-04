package routes

import (
	"net/http"
	"os"

	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/controllers"
	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/services"
	"github.com/eng-gabrielscardoso/goledger-challenge-besu/pkg/utils"
	"github.com/gin-gonic/gin"
)

var (
	transactionService    services.TransactionService       = *services.New()
	transactionController controllers.TransactionController = *controllers.New(&transactionService)
)

func SetupRouter() *gin.Engine {
	switch os.Getenv("GIN_MODE") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
		utils.Logger.Print("Running application in mode: ", os.Getenv("GIN_MODE"))
	case "test":
		gin.SetMode(gin.TestMode)
		utils.Logger.Print("Running application in mode: ", os.Getenv("GIN_MODE"))
	default:
		gin.SetMode(gin.DebugMode)
		utils.Logger.Print("Running application in mode: ", os.Getenv("GIN_MODE"))
	}

	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"message": "pong",
			},
		})
	})

	simpleStorage := router.Group("/simple-storage")
	{
		simpleStorage.GET("/", transactionController.GetValue)

		simpleStorage.POST("/", transactionController.SetValue)

		simpleStorage.POST("/sync", transactionController.SyncTransaction)

		simpleStorage.POST("/check", transactionController.CheckTransaction)
	}

	return router
}
