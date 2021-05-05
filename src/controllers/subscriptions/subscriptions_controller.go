package subscriptions

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chertokdmitry/surfapi/src/domain/subscriptions"
	"gitlab.com/chertokdmitry/surfapi/src/services"
	"gitlab.com/chertokdmitry/surfapi/src/utils/controllers_utils"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
	"net/http"
)

func GetByChatId(c *gin.Context) {
	id, errUtils := controllers_utils.CheckId(c.Param("chat_id"))
	if errUtils != nil {
		c.JSON(errUtils.Status, errUtils)
		return
	}
	spots, err := services.SubscriptionService.GetByChatId(id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, spots)
}

func Create(c *gin.Context) {
	var sub subscriptions.Sub

	if err := c.ShouldBindJSON(&sub); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	saveErr := services.SubscriptionService.Insert(&sub)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func Delete(c *gin.Context) {
	id, errUtils := controllers_utils.CheckId(c.Param("sub_id"))
	if errUtils != nil {
		c.JSON(errUtils.Status, errUtils)
		return
	}

	if err := services.SubscriptionService.Delete(id); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAll(c *gin.Context) {
	spots, err := services.SpotsService.GetAll()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, spots)
}
