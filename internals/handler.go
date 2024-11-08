package internals

import (
	"context"
	"go_backend/views"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

// This is basically a function that will render the html template to the user.
func render(ctx *gin.Context, status int, template templ.Component) error {
    ctx.Status(status)
    return template.Render(ctx.Request.Context(), ctx.Writer)
}

// This is a method that is attached to app config and returns a handlerFunc
func (app *Config)  HandleStockData() gin.HandlerFunc {
    return func(ctx *gin.Context) {

		// Basically creates a temperature request obj.
		temperatureRequest := &TemperatureRequest{
			Days: 7, 
		}

		// Get Temp data and err from GenerateTemperature function with request parameters sent.
		tempdata, err := app.GenerateTemperature(temperatureRequest)

		if err != nil {
            // If there's an error, respond with a 500 Internal Server Error status.
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

		// Basically returns the data in json format.
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

// Method attached to app with return type as gin.HandlerFunc.
func (app *Config) indexPageHandler() gin.HandlerFunc {
    return func(ctx *gin.Context) {
		// Creating a new context with timeout. This function is utilized to release any resources when the operation is complete.
        _, cancel := context.WithTimeout(context.Background(), appTimeout)

		// This defer basically scedules the cancel function to run at the end of the function. This is done to ensure that no more than 10 second takes to use this function.
        defer cancel()

		// Basically creates a temperature request obj.
		temperatureRequest := &TemperatureRequest{
            Days: 7, 
        }

		// Get Temp data and err from GenerateTemperature function with request parameters sent.
        response, err := app.GenerateTemperature(temperatureRequest)
        if err != nil {
			// Nil check if not nill show err in json format.
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

		// Preparing data in various format and assigning it in a variable.
		currentData := prepareCurrentData(response)
		dailyChartData := prepareDailyData(response)
		weeklyData := prepareWeeklyData(response)

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

