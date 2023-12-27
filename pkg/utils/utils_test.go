package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareCoordinats(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    *Geo
		wantErr bool
	}{
		{"Coordinats is valid", "211.433:-131.32", &Geo{Latitude: 211.433, Longitude: -131.32}, true},
		{"Coordinats is valid with coma", "211,433:-131,32", &Geo{Latitude: 211.433, Longitude: -131.32}, true},
		{"Coordinats is valid Latitude with coma", "211,433:-131.32", &Geo{Latitude: 211.433, Longitude: -131.32}, true},
		{"Coordinats is valid Longitude with coma", "211.433:-131,32", &Geo{Latitude: 211.433, Longitude: -131.32}, true},
		{"Coordinats is not valid", "211.433131", nil, false},
		{"Coordinats is not valid Longitude", "211.433:-131.32X", nil, false},
		{"Coordinats is not valid Latitude", "211.433X:-131.32", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrepareCoordinats(tt.text)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, err == nil, tt.wantErr)
		})
	}
}

func Test_SearchIndexAt(t *testing.T) {
	type args struct {
		haystack string
		needle   string
		offset   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Search from 0", args{"there any function I can use, where I can specify the start index", "can", 0}, 21},
		{"Search from 21", args{"there any function I can use, where I can specify the start index", "can", 20}, 21},
		{"Search from 25", args{"there any function I can use, where I can specify the start index", "can", 25}, 38},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchIndexAt(tt.args.haystack, tt.args.needle, tt.args.offset)
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_SearchWordIndexes(t *testing.T) {
	type args struct {
		text string
		word string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Text has same word multi entry 1", args{"asasasnanafnfsfs", "a"}, []int{0, 2, 4, 7, 9}},
		{"Text has same word multi entry 2", args{"there any function I can use, where I can specify the start index", "can"}, []int{21, 38}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchWordIndexes(tt.args.text, tt.args.word)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestFormatForecast(t *testing.T) {
	tests := []struct {
		name      string
		forecast  string
		wantCount int
	}{
		{"With word 'Day' for formatting", "Day: 2022-10-23", 2},
		{"With word 'Weather' for formatting", "Weather: clear sky", 1},
		{"Without word for formatting", "Weat her: clear sky", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatForecast(tt.forecast)
			assert.Equal(t, len(got), tt.wantCount)
		})
	}
}
