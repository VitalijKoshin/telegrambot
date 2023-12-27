package utils

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Words = []string{"Day", "Temperature", "Weather", "max:", "min:"}

type Geo struct {
	Latitude  float64
	Longitude float64
}

type MakeRequestI interface {
	Get(url string) (resp *http.Response, err error)
	ReadAll(r io.Reader) ([]byte, error)
}

type MakeRequestImpl struct{}

func TextHasCoordinats(text string) bool {
	_, err := PrepareCoordinats(text)

	return err == nil
}

func CommandHasCoordinats(command string, text string) bool {
	geoStr := strings.Trim(strings.Replace(text, "/"+command, "", 1), " ")
	_, err := PrepareCoordinats(geoStr)

	return err == nil
}

func CommandGetStrCoordinats(command string, text string) string {
	geoStr := strings.Trim(strings.Replace(text, "/"+command, "", 1), " ")
	_, err := PrepareCoordinats(geoStr)
	if err != nil {
		return ""
	}
	return geoStr
}

func PrepareCoordinats(text string) (*Geo, error) {
	text = strings.Replace(text, ",", ".", 2)
	data := strings.Split(text, ":")
	if len(data) != 2 {
		return nil, errors.New("please input \"latitude:longitude\" (Example: 33.44:-94.04)")
	}
	tmpLatitude, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		return nil, err
	}
	tmpLongitude, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		return nil, err
	}

	return &Geo{Latitude: tmpLatitude, Longitude: tmpLongitude}, nil
}

func SearchWordIndexes(text string, word string) []int {
	var result []int
	subStrIndex := 0
	for i := 0; i < len(text); i++ {
		subStrIndex = SearchIndexAt(text, word, subStrIndex)
		if subStrIndex == -1 {
			break
		}
		result = append(result, subStrIndex)
		subStrIndex++
	}
	return result
}

func FormatForecast(forecast string) []tgbotapi.MessageEntity {
	var entities []tgbotapi.MessageEntity
	for _, word := range Words {
		indxesWord := SearchWordIndexes(forecast, word)
		for _, indexWord := range indxesWord {
			switch word {
			case "Day":
				entities = append(entities, tgbotapi.MessageEntity{Type: "italic", Offset: indexWord, Length: len(word)})
				entities = append(entities, tgbotapi.MessageEntity{Type: "bold", Offset: indexWord, Length: len(word)})
			case "Weather":
				entities = append(entities, tgbotapi.MessageEntity{Type: "italic", Offset: indexWord, Length: len(word)})
			default:
				entities = append(entities, tgbotapi.MessageEntity{Type: "bold", Offset: indexWord, Length: len(word)})
			}
		}
	}

	return entities
}

func SearchIndexAt(haystack string, needle string, offset int) int {
	indexNeedle := strings.Index(haystack[offset:], needle)
	if indexNeedle > -1 {
		indexNeedle += offset
	}
	return indexNeedle
}

func (mr MakeRequestImpl) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

func (mr MakeRequestImpl) ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}
