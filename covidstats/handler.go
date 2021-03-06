package covidstats

import (
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

//Get Statewise Covid Stats from Geocoded latitude and longitude
func HandleStats(cache *redis.Client, db *mongo.Client) func(echo.Context) error {
	return func(ctx echo.Context) error {
		lat := ctx.QueryParam("lat")
		long := ctx.QueryParam("long")

		covidDataResponse, err := getLocationCovidStats(cache, db, lat, long)
		if err != nil {
			log.Println("Error getting covidDataResponse, err = ", err)
		}
		return ctx.JSON(http.StatusOK, covidDataResponse)
	}
}
