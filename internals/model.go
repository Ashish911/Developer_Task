package internals

import "time"

// WeatherCondition represents different weather types.
type WeatherCondition string

const (
	Sunny          WeatherCondition = "Sunny"
	PartlyCloudy   WeatherCondition = "Partly Cloudy"
	Cloudy         WeatherCondition = "Cloudy"
	Rainy          WeatherCondition = "Rainy"
)

// WindDirection represents cardinal directions.
type WindDirection string

const (
	North      WindDirection = "North"
	South      WindDirection = "South"
	East       WindDirection = "East"
	West       WindDirection = "West"
	NorthEast  WindDirection = "North-East"
	NorthWest  WindDirection = "North-West"
	SouthEast  WindDirection = "South-East"
	SouthWest  WindDirection = "South-West"
)

// TemperatureData holds the temperature and weather information for a specific time.
type TemperatureData struct {
	DateTime             time.Time         `json:"datetime"`
	Hour                 string            `json:"hour"` 
	Temperature          float64           `json:"temperature"`
	FeelsLike            float64           `json:"feels_like"`
	Weather              WeatherCondition   `json:"weather"`
	PrecipitationChance  float64           `json:"precipitation_chance"`
	WindSpeed            float64           `json:"wind_speed"`
	WindDirection        WindDirection      `json:"wind_direction"`
}

// TemperatureDataByDay represents temperature data organized by day.
type TemperatureDataByDay struct {
	Day    string                      `json:"day"`  
	Hours  []*TemperatureData    `json:"hours"`
}

// TemperatureRequest represents the parameters for requesting temperature data.
type TemperatureRequest struct {
	Days int `json:"days"`
}

// TemperatureResponse represents the API response for temperature data.
type TemperatureResponse struct {
	CurrentData  *TemperatureData             `json:"current_data"`
	DailyData    *TemperatureDataByDay        `json:"daily_data"`
	WeeklyData   map[string]*TemperatureDataByDay `json:"weekly_data"` 
}