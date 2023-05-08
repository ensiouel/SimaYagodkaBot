package owmutils

const (
	ClearSky        = `☀️`
	FewClouds       = `⛅`
	ScatteredClouds = `☁️`
	BrokenClouds
	ShowerRain   = `🌧️`
	Rain         = `🌦️`
	Thunderstorm = `🌩️`
	Snow         = `❄️`
	Mist         = `🌫️`
)

func GetIcon(id int) string {
	switch {
	case id >= 200 && id <= 232:
		return Thunderstorm
	case id >= 300 && id <= 300:
		return ShowerRain
	case id >= 500 && id <= 504:
		return Rain
	case id == 511:
		return Snow
	case id >= 520 && id <= 531:
		return ShowerRain
	case id >= 600 && id <= 622:
		return Snow
	case id >= 701 && id <= 781:
		return Mist
	case id == 800:
		return ClearSky
	case id == 801:
		return FewClouds
	case id == 802:
		return ScatteredClouds
	case id >= 803 && id <= 804:
		return BrokenClouds
	}

	return ""
}
