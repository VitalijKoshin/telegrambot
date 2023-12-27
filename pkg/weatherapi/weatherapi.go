package weatherapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"telegrambot/pkg/utils"

	"github.com/sirupsen/logrus"
)

const (
	SleepRequest     = 1
	UrlForecastByGeo = "https://api.openweathermap.org/data/3.0/onecall?appid=%s&lat=%f&lon=%f&exclude=%s&units=metric&lang=en"
)

type WeatherClient interface {
	GetWeatherForecastByLocation(geoUser GeoLocation) (string, error)
}

type weatherClientImpl struct {
	accessToken string
	mRequest    utils.MakeRequestI
}

type GeoLocation struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Exclude   string  `json:"exclude"`
}

type weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type forecastDaily struct {
	Timestamp   int                `json:"dt"`
	Temperature map[string]float64 `json:"temp"`
	FeelsLike   map[string]float64 `json:"feels_like"`
	Pressure    int                `json:"pressure"`
	Humidity    int                `json:"humidity"`
	WindSpeed   float64            `json:"wind_speed"`
	WindDeg     float64            `json:"wind_deg"`
	WindGust    float64            `json:"wind_gust"`
	Clouds      int                `json:"clouds"`
	Pop         float64            `json:"pop"`
	Weather     []weather          `json:"weather"`
}

type Forecast struct {
	Latitude  float64         `json:"lat"`
	Longitude float64         `json:"lon"`
	Timezone  string          `json:"timezone"`
	Daily     []forecastDaily `json:"daily"`
}

var mut sync.Mutex

func CreateClient(accessToken string) WeatherClient {
	client := &weatherClientImpl{
		accessToken: accessToken,
		mRequest:    utils.MakeRequestImpl{},
	}
	return client
}

func (c *weatherClientImpl) GetWeatherForecastByLocation(geoUser GeoLocation) (string, error) {
	url := fmt.Sprintf(UrlForecastByGeo,
		c.accessToken,
		geoUser.Latitude,
		geoUser.Longitude,
		geoUser.Exclude)
	body, err := c.makeRequest(url)
	if err != nil {
		return "", err
	}
	return c.weatherForecastRender(body)
}

func (c *weatherClientImpl) weatherForecastRender(body []byte) (string, error) {
	var forecast Forecast
	result := ""
	err := json.Unmarshal(body, &forecast)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("Latitude: %f\nLongitude: %f\nTimezone: %s\n", forecast.Latitude, forecast.Longitude, forecast.Timezone)

	for _, value := range forecast.Daily {
		day := time.Unix(int64(value.Timestamp), 0)
		result += fmt.Sprintf("\nDay: %s\nWeather:", day.Format("2006-01-02"))
		for _, wv := range value.Weather {
			result += fmt.Sprintf("\n%s", wv.Description)
		}
		result += "\nTemperature:"
		for tk, tv := range value.Temperature {
			result += fmt.Sprintf("\n%s:		%d deg C", tk, int(tv))
		}
		result += "\n"
		break
	}

	return result, nil
}

func (c *weatherClientImpl) makeRequest(url string) ([]byte, error) {
	mut.Lock()
	defer mut.Unlock()
	time.Sleep(SleepRequest * time.Second)
	resp, err := c.mRequest.Get(url)
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := c.mRequest.ReadAll(resp.Body)
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("makeRequest status is not ok (" + resp.Status + ")")
		logrus.Debug(err)
		return nil, err
	}

	return body, err
}
