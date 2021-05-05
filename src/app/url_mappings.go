package app

import (
	"gitlab.com/chertokdmitry/surfapi/src/controllers/cameras"
	"gitlab.com/chertokdmitry/surfapi/src/controllers/images"
	"gitlab.com/chertokdmitry/surfapi/src/controllers/ping"
	"gitlab.com/chertokdmitry/surfapi/src/controllers/spots"
	"gitlab.com/chertokdmitry/surfapi/src/controllers/subscriptions"
	"gitlab.com/chertokdmitry/surfapi/src/controllers/users"
	"gitlab.com/chertokdmitry/surfapi/src/controllers/weather"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/weather/:spot_id", weather.Get)

	router.GET("/spots/region/:region_id", spots.GetByRegionId)
	router.GET("/spots/all", spots.GetAll)

	router.GET("/cameras/all", cameras.GetAll)
	router.GET("/cameras/spot/:spot_id", cameras.GetBySpot)

	router.POST("images", images.Create)

	router.GET("/subscriptions/chat/:chat_id", subscriptions.GetByChatId)
	router.POST("subscriptions", subscriptions.Create)
	router.DELETE("/subscriptions/:sub_id", subscriptions.Delete)

	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
