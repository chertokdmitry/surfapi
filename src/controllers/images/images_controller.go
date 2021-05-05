package images

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chertokdmitry/surfapi/src/domain/images"
	"gitlab.com/chertokdmitry/surfapi/src/services"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
	"net/http"
)

func Create(c *gin.Context) {
	var image images.Image

	if err := c.ShouldBindJSON(&image); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	saveErr := services.ImageService.Insert(&image)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}
