package message

const (
	ErrOpenConn      string = "error when trying to open postgres connection"
	ErrPingConn      string = "error when ping postgres connection"
	ErrQueryGetSpot  string = "error query to get all spots"
	ErrQueryGetSubs  string = "error query to get subs"
	ErrScanSpots     string = "error scan spots"
	ErrScanSubs      string = "error scan subs"
	Russia           string = "\xE2\x98\xAD Россия"
	Worldwide        string = "\xE2\x93\xA6 Мир"
	Subscribe        string = "\xE2\x88\xBB Подписаться"
	SubList          string = "\xE2\x89\x8B Мои подписки"
	ErrSendInline    string = "error when send inline message"
	Back             string = "Вернуться"
	ErrUnmarshal     string = "error when marshal db_resource"
	ErrInsertSub     string = "error when insert new sub"
	ErrDeleteSub     string = "error when delete sub"
	ErrSelectWeather string = "error when select weather by spot"
	ErrZeroRows      string = "error, zero rows found"
	ErrUpdateMess    string = "error when update messages"
	Forecast24       string = "Прогноз ветра на 24 часа."
	SubMessage       string = "Введите номер спота для подписки:"
	RussiaRegionId   int    = 190
	WorldRegionId    int    = 290
	GetHour          string = "Выберите время отправки"
	GetSpot          string = "Выберите спот"
	WrongSpotId      string = "Введите правильный номер спота"
	SubSuccess       string = "Вы успешно подписались!"
	SubDelSuccess    string = "Подписка удалена"
	Delete           string = " \xE2\x98\x92 удалить"
)
