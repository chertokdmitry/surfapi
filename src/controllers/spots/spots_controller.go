package spots

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chertokdmitry/surfapi/src/services"
	"gitlab.com/chertokdmitry/surfapi/src/utils/controllers_utils"
	"net/http"
)

func GetByRegionId(c *gin.Context) {
	id, errUtils := controllers_utils.CheckId(c.Param("region_id"))
	if errUtils != nil {
		c.JSON(errUtils.Status, errUtils)
		return
	}
	spots, err := services.SpotsService.GetByRegionId(id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, spots)
}

func GetAll(c *gin.Context) {
	spots, err := services.SpotsService.GetAll()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, spots)
}
