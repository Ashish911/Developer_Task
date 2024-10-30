package internals

import (
	"strconv"
	"time"

	"math/rand"
)

func GenerateConditions(temperature float64) (WeatherCondition, float64) {
	var condition WeatherCondition
	var chance float64

	switch {
	case temperature >= 20:
		condition = Sunny
		chance = 10.0 
	case temperature >= 15:
		condition = PartlyCloudy
		chance = 20.0
	case temperature >= 10:
		condition = Cloudy
		chance = 30.0
	default:
		condition = Rainy
		chance = 80.0
	}

	return condition, chance
}

func GenerateWindConditions() (float64, WindDirection) {
	// Here i decided to put the wind speed between 0 and 20 because almost always it is like this in UK. But not lately, but i know it will happen in a few weeks.
	windSpeed := rand.Float64() * 20

	// Randomly set a wind direction.
	directions := []WindDirection{North, South, East, West, NorthEast, NorthWest, SouthEast, SouthWest}
	direction := directions[rand.Intn(len(directions))]

	return windSpeed, direction
}

func CalculateFeelsLikeTemperature(temperature, windSpeed float64) float64 {
	// This is a simple formula to generate feels like temp found in the internet.
	if windSpeed > 0 {
		return temperature - (windSpeed / 3.0)
	}

	return temperature
}

func (app *Config) GenerateTemperature(req *TemperatureRequest) (*TemperatureResponse, error) {
	weeklyData := make(map[string]*TemperatureDataByDay)  
    now := time.Now()
    var currentData *TemperatureData                      
    var dailyData *TemperatureDataByDay 

	// Generate data for each day
	for dayOffset := 0; dayOffset < req.Days; dayOffset++ {
        day := now.AddDate(0, 0, dayOffset).Format("2006-01-02")  // Format as "YYYY-MM-DD"
        dayData := &TemperatureDataByDay{
            Day:   day,
            Hours: make([]*TemperatureData, 24),
        }

        for hour := 0; hour < 24; hour++ {
            
            // Simulate temperature and weather fluctuations
            currentTemp := 15.0 + rand.Float64()*10.0
            fluctuation := (rand.Float64() - 0.5) * 2
            currentTemp += fluctuation

            // Keep temperature within a reasonable range
            if currentTemp < 5 {
                currentTemp = 5
            } else if currentTemp > 35 {
                currentTemp = 35
            }

            // Generate other weather data
            windSpeed, windDirection := GenerateWindConditions()
            weather, precipitationChance := GenerateConditions(currentTemp)
            feelsLikeTemp := CalculateFeelsLikeTemperature(currentTemp, windSpeed)

            // Create the hourly temperature data
            tempData := &TemperatureData{
                Hour: strconv.Itoa(hour),
                DateTime:            time.Date(now.Year(), now.Month(), now.Day()+dayOffset, hour, 0, 0, 0, time.UTC),
                Temperature:         currentTemp,
                FeelsLike:           feelsLikeTemp,
                Weather:             weather,
                PrecipitationChance: precipitationChance,
                WindSpeed:           windSpeed,
                WindDirection:       windDirection,
            }

            // Add this hour's data to the day's map
            dayData.Hours[hour] = tempData

            // Set current data if it's today's date and the current hour
            if dayOffset == 0 && hour == now.Hour() {
                currentData = tempData
            }
        }

        // Add today's hourly data to the weekly data map
        weeklyData[day] = dayData

        // Set daily data if it’s today
        if dayOffset == 0 {
            dailyData = dayData
        }
    }

	response := &TemperatureResponse{
        CurrentData: currentData,
        DailyData:   dailyData,
        WeeklyData:  weeklyData,
    }

    return response, nil
}