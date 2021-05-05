package weather

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chertokdmitry/surfapi/src/services"
	"gitlab.com/chertokdmitry/surfapi/src/utils/controllers_utils"
	"net/http"
)

func Get(c *gin.Context) {
	spotId, err := controllers_utils.CheckId(c.Param("spot_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	weather, err := services.WeatherService.GetWeather(spotId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, weather)
}
