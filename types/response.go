package types

type ApiResponse struct {
	Coord      coord     `json:"coord"`
	Weather    []weather `json:"weather"`
	Base       string    `json:"base"`
	Main       main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       wind      `json:"wind"`
	Rain       rain      `json:"rain"`
	Clouds     clouds    `json:"clouds"`
	DT         int64     `json:"dt"`
	Sys        sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

type coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type clouds struct {
	All int `json:"all"`
}

type rain struct {
	OneH float64 `json:"1h"`
}

type sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}
