package internals

import (
	"go_backend/views"
	"sort"
	"time"
)

// prepareCurrentData processes and returns current temperature data
func prepareCurrentData(response *TemperatureResponse) *views.TemperatureData {
    if response.CurrentData == nil {
        return nil
    }
    return &views.TemperatureData{
        DateTime:            response.CurrentData.DateTime,
        Hour:                response.CurrentData.Hour,
        Temperature:         response.CurrentData.Temperature,
        FeelsLike:           response.CurrentData.FeelsLike,
        Weather:             string(response.CurrentData.Weather),
        PrecipitationChance: response.CurrentData.PrecipitationChance,
        WindSpeed:           response.CurrentData.WindSpeed,
        WindDirection:       string(response.CurrentData.WindDirection),
    }
}

// prepareDailyData prepares daily chart data from the response
func prepareDailyData(response *TemperatureResponse) *views.DailyChartData {
    dailyChartData := &views.DailyChartData{
        Labels:         []string{},
        TemperatureData: []float64{},
        FeelsLikeData:  []float64{},
    }
    for _, data := range response.DailyData.Hours {
        dailyChartData.Labels = append(dailyChartData.Labels, data.Hour)
        dailyChartData.TemperatureData = append(dailyChartData.TemperatureData, data.Temperature)
        dailyChartData.FeelsLikeData = append(dailyChartData.FeelsLikeData, data.FeelsLike)
    }
    return dailyChartData
}

// prepareWeeklyData aggregates weekly data with min, max, average calculations
func prepareWeeklyData(response *TemperatureResponse) []*views.WeeklyChartData {
    tomorrow := time.Now().Add(24 * time.Hour).Format("2006-01-02")
    weeklyData := []*views.WeeklyChartData{}

    for day, dayData := range response.WeeklyData {
        if day < tomorrow {
            continue
        }

        var minTemp, maxTemp, tempSum, windSpeedSum, precipSum float64
        count := len(dayData.Hours)

        if count == 0 {
            continue
        }

        minTemp = dayData.Hours[0].Temperature
        maxTemp = dayData.Hours[0].Temperature

        for _, hourData := range dayData.Hours {
            temp := hourData.Temperature
            windSpeed := hourData.WindSpeed
            precipChance := hourData.PrecipitationChance

            if temp < minTemp {
                minTemp = temp
            }
            if temp > maxTemp {
                maxTemp = temp
            }

            tempSum += temp
            windSpeedSum += windSpeed
            precipSum += float64(precipChance)
        }

        avgTemp := tempSum / float64(count)
        avgWindSpeed := windSpeedSum / float64(count)
        avgPrecipChance := precipSum / float64(count)

        dayFormatted, _ := time.Parse("2006-01-02", day)
        formattedDay := dayFormatted.Format("02 Jan 2006")

        weeklyData = append(weeklyData, &views.WeeklyChartData{
            Day:                   formattedDay,
            MinTemp:               minTemp,
            MaxTemp:               maxTemp,
            AvgTemp:               avgTemp,
            AvgWindSpeed:          avgWindSpeed,
            AvgPrecipitationChance: avgPrecipChance,
        })
    }

    sort.Slice(weeklyData, func(i, j int) bool {
        day1, _ := time.Parse("02 Jan 2006", weeklyData[i].Day)
        day2, _ := time.Parse("02 Jan 2006", weeklyData[j].Day)
        return day1.Before(day2)
    })

    return weeklyData
}