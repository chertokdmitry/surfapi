package app

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chertokdmitry/surfapi/src/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start app")
	router.Run(":8080")

}
