package cameras

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chertokdmitry/surfapi/src/services"
	"gitlab.com/chertokdmitry/surfapi/src/utils/controllers_utils"
	"net/http"
)

func GetAll(c *gin.Context) {
	cameras, err := services.CamerasService.GetAll()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, cameras)
}

func GetBySpot(c *gin.Context) {
	id, errUtils := controllers_utils.CheckId(c.Param("spot_id"))
	if errUtils != nil {
		c.JSON(errUtils.Status, errUtils)
		return
	}

	cameras, err := services.CamerasService.GetBySpot(id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, cameras)
}
