package internals

import (
	"context"
	"fmt"
	"go_backend/views"
	"net/http"
	"sort"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func render(ctx *gin.Context, status int, template templ.Component) error {
    ctx.Status(status)
    return template.Render(ctx.Request.Context(), ctx.Writer)
}

func (app *Config)  HandleStockData() gin.HandlerFunc {
    return func(ctx *gin.Context) {

		temperatureRequest := &TemperatureRequest{
			Days: 7, 
		}

		tempdata, err := app.GenerateTemperature(temperatureRequest)

		if err != nil {
            // If there's an error, respond with a 500 Internal Server Error status.
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

		ctx.JSON(http.StatusOK, tempdata)
	}
}

// Keeping these code for study purposes.
// func (app *Config) oldIndexPageHandler() gin.HandlerFunc {
//     return func(ctx *gin.Context) {
//         _, cancel := context.WithTimeout(context.Background(), appTimeout)
//         defer cancel()

// 		temperatureRequest := &TemperatureRequest{
//             Days: 7, 
//         }

//         response, err := app.GenerateTemperature(temperatureRequest)
//         if err != nil {
//             ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//             return
//         }

// 		var currentData *views.TemperatureData
//         if response.CurrentData != nil {
// 			current := &views.TemperatureData{
// 				DateTime: response.CurrentData.DateTime,
// 				Hour: response.CurrentData.Hour,
// 				Temperature: response.CurrentData.Temperature,
// 				FeelsLike: response.CurrentData.FeelsLike,
// 				Weather: string(response.CurrentData.Weather),
// 				PrecipitationChance: response.CurrentData.PrecipitationChance,
// 				WindSpeed: response.CurrentData.WindSpeed,
// 				WindDirection: string(response.CurrentData.WindDirection),
// 			} // Safely dereference if not nil

// 			currentData = current
//         }



//         // Declare the slice to hold pointers to TemperatureDataByDay
// 		// var dailyData *views.TemperatureDataByDay

// 		// dailyDataMap := make(map[string]map[int]*views.TemperatureData)


// 		//  // Populate the dailyDataMap from the response
//         // for _, hourData := range response.DailyData.Hours {
//         //     day := hourData.DateTime.Format("2006-01-02") // Get the day in YYYY-MM-DD format
//         //     hour := hourData.DateTime.Hour()               // Get hour as an integer

//         //     // Initialize the map for the day if it doesn't exist
//         //     if dailyDataMap[day] == nil {
//         //         dailyDataMap[day] = make(map[int]*views.TemperatureData)
//         //     }

//         //     // Create a new TemperatureData instance
//         //     tempData := &views.TemperatureData{
//         //         DateTime:             hourData.DateTime,
//         //         Hour:                 hourData.Hour,
//         //         Temperature:          hourData.Temperature,
//         //         FeelsLike:            hourData.FeelsLike,
//         //         Weather:              string(hourData.Weather),
//         //         PrecipitationChance:  hourData.PrecipitationChance,
//         //         WindSpeed:            hourData.WindSpeed,
//         //         WindDirection:        string(hourData.WindDirection),
//         //     }

//         //     // Assign the TemperatureData to the appropriate hour in the map
//         //     dailyDataMap[day][hour] = tempData // Use the pointer to views.TemperatureData
//         // }

// 		// for day, hours := range dailyDataMap {
//         //     dailyData = &views.TemperatureDataByDay{
//         //         Day:   day,
//         //         Hours: hours, // Use the map[int]*views.TemperatureData
//         //     }
//         // }

// 		dailyChartData := &views.DailyChartData{
//             Labels:         []string{},
//             TemperatureData: []float64{},
//             FeelsLikeData:  []float64{},
//         }

// 		for _, data := range response.DailyData.Hours {
//             dailyChartData.Labels = append(dailyChartData.Labels, data.Hour) // Or format data.DateTime
//             dailyChartData.TemperatureData = append(dailyChartData.TemperatureData, data.Temperature)
//             dailyChartData.FeelsLikeData = append(dailyChartData.FeelsLikeData, data.FeelsLike)
//         }
		

// 		// Assemble the data into a view model
//         viewModel := views.TemperatureDataViewModel{
//             CurrentData: currentData,
// 			DailyData: dailyChartData,
//             // DailyData:   dailyData,
//             // WeeklyData:  make(map[string]*views.TemperatureDataByDay),
//         }

// 		// Populate the WeeklyData map
// 		// for k, v := range response.WeeklyData {
// 		// 	viewModel.WeeklyData[k] = &views.TemperatureDataByDay{
// 		// 		Day:   v.Day,
// 		// 		Hours: convertTemperatureData(v.Hours),
// 		// 	}
// 		// }
	
// 		// Render with the view model
// 		if err := render(ctx, http.StatusOK, views.Index(viewModel)); err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to render template"})
// 			return
// 		}
//     }
// }

// func convertTemperatureData(data []*TemperatureData) []*views.TemperatureData {
// 	result := make([]*views.TemperatureData, 24)
// 	for k, v := range data {
// 		result[k] = &views.TemperatureData{
// 			DateTime:                 v.DateTime,
// 			Hour:                v.Hour,
// 			Temperature:         v.Temperature,
// 			FeelsLike:           v.FeelsLike,
// 			Weather:             string(v.Weather),
// 			PrecipitationChance: v.PrecipitationChance,
// 			WindSpeed:           v.WindSpeed,
// 			WindDirection:       string(v.WindDirection),
// 		}
// 	}
// 	return result
// }

func (app *Config) indexPageHandler() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        _, cancel := context.WithTimeout(context.Background(), appTimeout)
        defer cancel()

		temperatureRequest := &TemperatureRequest{
            Days: 7, 
        }

        response, err := app.GenerateTemperature(temperatureRequest)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

		var currentData *views.TemperatureData
        if response.CurrentData != nil {
			current := &views.TemperatureData{
				DateTime: response.CurrentData.DateTime,
				Hour: response.CurrentData.Hour,
				Temperature: response.CurrentData.Temperature,
				FeelsLike: response.CurrentData.FeelsLike,
				Weather: string(response.CurrentData.Weather),
				PrecipitationChance: response.CurrentData.PrecipitationChance,
				WindSpeed: response.CurrentData.WindSpeed,
				WindDirection: string(response.CurrentData.WindDirection),
			} // Safely dereference if not nil

			currentData = current
        }

		dailyChartData := &views.DailyChartData{
            Labels:         []string{},
            TemperatureData: []float64{},
            FeelsLikeData:  []float64{},
        }

		for _, data := range response.DailyData.Hours {
            dailyChartData.Labels = append(dailyChartData.Labels, data.Hour) // Or format data.DateTime
            dailyChartData.TemperatureData = append(dailyChartData.TemperatureData, data.Temperature)
            dailyChartData.FeelsLikeData = append(dailyChartData.FeelsLikeData, data.FeelsLike)
        }

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

                // Update min and max temperature
                if temp < minTemp {
                    minTemp = temp
                }
                if temp > maxTemp {
                    maxTemp = temp
                }

                // Accumulate for averages
                tempSum += temp
                windSpeedSum += windSpeed
                precipSum += float64(precipChance)
            }

            avgTemp := tempSum / float64(count)
            avgWindSpeed := windSpeedSum / float64(count)
            avgPrecipChance := precipSum / float64(count)

            // Format day for display, e.g., "30 Oct 2024"
            dayFormatted, _ := time.Parse("2006-01-02", day)
            formattedDay := dayFormatted.Format("02 Jan 2006")

            dailySummary := &views.WeeklyChartData{
                Day:                   formattedDay,
                MinTemp:               minTemp,
                MaxTemp:               maxTemp,
                AvgTemp:               avgTemp,
                AvgWindSpeed:          avgWindSpeed,
                AvgPrecipitationChance: avgPrecipChance,
            }

            weeklyData = append(weeklyData, dailySummary)
        }

		sort.Slice(weeklyData, func(i, j int) bool {
			day1, err1 := time.Parse("02 Jan 2006", weeklyData[i].Day)
			day2, err2 := time.Parse("02 Jan 2006", weeklyData[j].Day)
		
			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing date for sorting:", err1, err2)
				return false
			}
		
			return day1.Before(day2)
		})

		// Assemble the data into a view model
        viewModel := views.TemperatureDataViewModel{
            CurrentData: currentData,
			DailyData: dailyChartData,
            WeeklyData: weeklyData,
        }
	
		// Render with the view model
		if err := render(ctx, http.StatusOK, views.Index(viewModel)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to render template"})
			return
		}
    }
}

