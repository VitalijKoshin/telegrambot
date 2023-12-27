package services

import (
	"testing"

	"telegrambot/models"
	"telegrambot/repository"

	"github.com/stretchr/testify/assert"
)

func TestUserService_SubscribeUser(t *testing.T) {
	tests := []struct {
		name           string
		userRepository repository.UserRepository
		user           *models.User
		want           bool
		wantErr        bool
	}{
		{"Usual user subscribe", &FakeUserModel{}, &models.User{TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, true, false},
		{"Usual user subscribed", &FakeUserModelSubscribed{}, &models.User{TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, true, false},
		{"Error subscribe", &FakeUserModelError{}, &models.User{TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserServiceImpl{
				userRepository: tt.userRepository,
			}
			got, err := u.SubscribeUser(tt.user)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, (err != nil), tt.wantErr)
		})
	}
}

func TestUserService_UnSubscribeUser(t *testing.T) {
	tests := []struct {
		name           string
		userRepository repository.UserRepository
		user           *models.User
		want           bool
		wantErr        bool
	}{
		{"Usual user unsubscribe", &FakeUserModel{}, &models.User{TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, true, false},
		{"Usual user unsubscribed", &FakeUserModelSubscribed{}, &models.User{TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, true, false},
		{"Error unsubscribe", &FakeUserModelError{}, &models.User{TID: 9223372036854775806, Username: "U", FirstName: "F", LastName: "L", Status: true}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserServiceImpl{
				userRepository: tt.userRepository,
			}
			got, err := u.UnSubscribeUser(tt.user)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, (err != nil), tt.wantErr)
		})
	}
}
