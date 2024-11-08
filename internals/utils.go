package internals

import (
	"go_backend/views"
	"sort"
	"time"
)

/*
   So in simple terms this function prepareCurrentData processes and returns current temperature data.
   Basically what this function is doing is taking a response pointer named as TemperatureResponse and returning a TemperatureData pointer from the views section of our code.
   First thing done here is nil checking, as it is a crucial thing in any programming to always look for null pointer exception. Apparently it seems like normal struct copying
   without pointer does not allow for returning nil. Here in the &views what it is doing is creating an instance of the return type by adding & it shows that the return type is a
   pointer, inside our Golang will automatically dereference the pointer values and put it in the new pointer.
*/
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

/*
    PrepareDailyData prepares daily chart data from the response.
    Basically iterates over the response which is based on TemperatureResponse and appends the hour, temp and feelslike temp data in each empty struct from above. A blank identifier is 
    used here i.e _ because index was not required for this operation.
*/
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

/*
    PrepareWeeklyData aggregates weekly data with min, max, average calculations.

*/

func prepareWeeklyData(response *TemperatureResponse) []*views.WeeklyChartData {
    // This is basically generating or calculating tomorrows date.
    tomorrow := time.Now().Add(24 * time.Hour).Format("2006-01-02")

    // Initializing and declaring variable based on struct.
    weeklyData := []*views.WeeklyChartData{}

    for day, dayData := range response.WeeklyData {
        // Basically going through our response and calculating data accordfing to each days.
        // This like basically checks for if the current day is less than tomorrow meaning it is going to be currentdata. Since our objective is to get the next 6 days worth of data removing todays data.
        if day < tomorrow {
            continue
        }

        var minTemp, maxTemp, tempSum, windSpeedSum, precipSum float64
        count := len(dayData.Hours)

        if count == 0 {
            continue
        }

        // Initializing the mintemp and the max temp to be the first temp of the iteration.
        minTemp = dayData.Hours[0].Temperature
        maxTemp = dayData.Hours[0].Temperature

        for _, hourData := range dayData.Hours {
            temp := hourData.Temperature
            windSpeed := hourData.WindSpeed
            precipChance := hourData.PrecipitationChance

            // Checking current temps, and based on the current temp assigning mintemp and maxtemp.
            if temp < minTemp {
                minTemp = temp
            }
            if temp > maxTemp {
                maxTemp = temp
            }

            // Summating all the temps, windspeed and precipitation for avg calculation.
            tempSum += temp
            windSpeedSum += windSpeed
            precipSum += float64(precipChance)
        }

        // Calculating the avg of all the required vars or fields.
        avgTemp := tempSum / float64(count)
        avgWindSpeed := windSpeedSum / float64(count)
        avgPrecipChance := precipSum / float64(count)


        dayFormatted, _ := time.Parse("2006-01-02", day)
        formattedDay := dayFormatted.Format("02 Jan 2006")

        // Here basically creating a copy of struct and appending value in it.
        weeklyData = append(weeklyData, &views.WeeklyChartData{
            Day:                   formattedDay,
            MinTemp:               minTemp,
            MaxTemp:               maxTemp,
            AvgTemp:               avgTemp,
            AvgWindSpeed:          avgWindSpeed,
            AvgPrecipitationChance: avgPrecipChance,
        })
    }

    // Basically sorts the weekly data by day in ascending order before returning it.
    sort.Slice(weeklyData, func(i, j int) bool {
        day1, _ := time.Parse("02 Jan 2006", weeklyData[i].Day)
        day2, _ := time.Parse("02 Jan 2006", weeklyData[j].Day)
        return day1.Before(day2)
    })

    return weeklyData
}