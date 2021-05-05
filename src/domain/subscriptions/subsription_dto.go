package subscriptions

var IsSpotId, IsHour bool
var SpotId, Hour int

type Sub struct {
	Id     int64  `json:"id"`
	ChatId int64  `json:"chat_id"`
	SpotId int64  `json:"spot_id"`
	Title  string `json:"title"`
	Hour   string `json:"hour"`
}
