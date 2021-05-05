package weather

type Weather struct {
	SpotId int64  `json:"spot_id"`
	Title  string `json:"title"`
	Hourly string `json:"hourly"`
}

type Hour struct {
	Clouds     int     `json:"clouds"`
	Wind_deg   int     `json:"wind_deg"`
	Wind_speed float32 `json:"wind_speed"`
	Temp       float32 `json:"temp"`
}
