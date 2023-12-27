package weatherapi

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"telegrambot/pkg/utils"

	"github.com/stretchr/testify/assert"
)

type MakeRequestTestImpl struct{}

func (mr MakeRequestTestImpl) Get(url string) (resp *http.Response, err error) {
	return &http.Response{Body: http.NoBody, StatusCode: http.StatusOK}, nil
}

func (mr MakeRequestTestImpl) ReadAll(r io.Reader) ([]byte, error) {
	return []byte(``), nil
}

type MakeRequestTestImplErrorGet struct{}

func (mr MakeRequestTestImplErrorGet) Get(url string) (resp *http.Response, err error) {
	return nil, fmt.Errorf("url Get error")
}

func (mr MakeRequestTestImplErrorGet) ReadAll(r io.Reader) ([]byte, error) {
	return []byte(``), nil
}

type MakeRequestTestImplErrorReadAll struct{}

func (mr MakeRequestTestImplErrorReadAll) Get(url string) (resp *http.Response, err error) {
	return &http.Response{Body: http.NoBody, StatusCode: http.StatusOK}, nil
}

func (mr MakeRequestTestImplErrorReadAll) ReadAll(r io.Reader) ([]byte, error) {
	return []byte(nil), fmt.Errorf("url ReadAll error")
}

type MakeRequestTestRespImpl struct{}

func (mr MakeRequestTestRespImpl) Get(url string) (resp *http.Response, err error) {
	return &http.Response{Body: http.NoBody, StatusCode: http.StatusOK}, nil
}

func (mr MakeRequestTestRespImpl) ReadAll(r io.Reader) ([]byte, error) {
	return []byte(stringReadAll), nil
}

var stringReadAll = `{
	"lat": 39.31,
	"lon": -74.5,
	"timezone": "America/New_York",
	"timezone_offset": -18000,
	"current": {
	  "dt": 1646318698,
	  "sunrise": 1646306882,
	  "sunset": 1646347929,
	  "temp": 282.21,
	  "feels_like": 278.41,
	  "pressure": 1014,
	  "humidity": 65,
	  "dew_point": 275.99,
	  "uvi": 2.55,
	  "clouds": 40,
	  "visibility": 10000,
	  "wind_speed": 8.75,
	  "wind_deg": 360,
	  "wind_gust": 13.89,
	  "weather": [
		{
		  "id": 802,
		  "main": "Clouds",
		  "description": "scattered clouds",
		  "icon": "03d"
		}
	  ]
	}
}`

func Test_weatherClientImpl_GetWeatherForecastByLocation(t *testing.T) {
	tests := []struct {
		name        string
		accessToken string
		geoUser     GeoLocation
		mRequest    utils.MakeRequestI
		wantError   error
	}{
		{
			"Normal forecast for geolocation",
			"bee5ff753b73cef9cc381d76940991ae",
			GeoLocation{Latitude: 43.2316846, Longitude: 27.9730744, Exclude: "current"},
			&MakeRequestTestRespImpl{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weatherClient := weatherClientImpl{
				accessToken: tt.accessToken,
				mRequest:    tt.mRequest,
			}
			_, err := weatherClient.GetWeatherForecastByLocation(tt.geoUser)

			assert.Equal(t, err, tt.wantError)
		})
	}
}

func Test_weatherClientImpl_makeRequest(t *testing.T) {
	type fields struct {
		accessToken string
		mRequest    utils.MakeRequestI
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		err    error
	}{
		{
			"Normal empty make request",
			fields{accessToken: "accessToken", mRequest: &MakeRequestTestImpl{}},
			args{url: "https://testurl.net"},
			[]byte(``),
			nil,
		},
		{
			"Normal make request",
			fields{accessToken: "accessToken", mRequest: &MakeRequestTestRespImpl{}},
			args{url: "https://testurl.net"},
			[]byte(stringReadAll),
			nil,
		},
		{
			"Error make request Get ",
			fields{accessToken: "accessToken", mRequest: &MakeRequestTestImplErrorGet{}},
			args{url: "https://testurl.net"},
			[]byte(nil),
			fmt.Errorf("url Get error"),
		},
		{
			"Error make request ReadAll ",
			fields{accessToken: "accessToken", mRequest: &MakeRequestTestImplErrorReadAll{}},
			args{url: "https://testurl.net"},
			[]byte(nil),
			fmt.Errorf("url ReadAll error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			c := &weatherClientImpl{
				accessToken: tt.fields.accessToken,
				mRequest:    tt.fields.mRequest,
			}
			got, err := c.makeRequest(tt.args.url)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, err, tt.err)
			end := time.Now()
			assert.GreaterOrEqual(t, end.Sub(start), (SleepRequest * time.Second))
		})
	}
}
