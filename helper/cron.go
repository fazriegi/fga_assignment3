package helper

import (
	"math/rand"

	"github.com/fazriegi/fga_assignment3/model"
	"github.com/robfig/cron"
)

func StartCron() {
	Weather := model.Weather{}

	c := cron.New()
	c.AddFunc("*/15 * * * * *", func() {
		Weather.Status.Wind = rand.Intn(100) + 1
		Weather.Status.Water = rand.Intn(100) + 1

		WriteFile(Weather, "weather.json")
	})
	c.Start()
}
