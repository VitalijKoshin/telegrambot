package services

import (
	"testing"

	"telegrambot/models"
	"telegrambot/repository"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSubsService_CreateSubs(t *testing.T) {
	subsModelFakeOk := FakeSubsModel{}
	type fields struct {
		subsRepository repository.SubsRepository
	}
	type args struct {
		chatID    int64
		cTime     int64
		uid       primitive.ObjectID
		latitude  float64
		longitude float64
		frequency int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   models.Subs
	}{
		{
			"Usual create subs",
			fields{subsRepository: &subsModelFakeOk},
			args{cTime: currentTime, uid: testSubsUID, latitude: 22.00, longitude: 99.00, frequency: 86400},
			models.Subs{
				ID:        testSubsID,
				UID:       testSubsUID,
				CTime:     currentTime,
				UTime:     currentTime,
				LTime:     0,
				NTime:     currentTime + 86400,
				Latitude:  22.00,
				Longitude: 99.00,
				Frequency: 86400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubsServiceImpl{
				subsRepository: tt.fields.subsRepository,
			}
			got := s.CreateSubs(tt.args.chatID, tt.args.cTime, tt.args.uid, tt.args.latitude, tt.args.longitude, tt.args.frequency)
			got.ID = testSubsID
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestSubsService_SaveSubs(t *testing.T) {
	subsModelFakeOk := FakeSubsModel{}
	subsModelFakeErr := FakeSubsErrorModel{}

	type fields struct {
		subsRepository repository.SubsRepository
	}
	type args struct {
		subs *models.Subs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Subs
		wantErr bool
	}{
		{"Ususal save subs", fields{&subsModelFakeOk}, args{&models.Subs{}}, &models.Subs{}, false},
		{"Error save subs", fields{&subsModelFakeErr}, args{&models.Subs{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubsServiceImpl{
				subsRepository: tt.fields.subsRepository,
			}
			got, err := s.SaveSubs(tt.args.subs)
			assert.Equal(t, err != nil, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestSubsService_DeleteSubs(t *testing.T) {
	subsModelFakeOk := FakeSubsModel{}
	subsModelFakeErr := FakeSubsErrorModel{}
	type fields struct {
		subsRepository repository.SubsRepository
	}
	type args struct {
		subs *models.Subs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{"Ususal delete subs", fields{&subsModelFakeOk}, args{&models.Subs{}}, true, false},
		{"Error delete subs", fields{&subsModelFakeErr}, args{&models.Subs{}}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubsServiceImpl{
				subsRepository: tt.fields.subsRepository,
			}
			got, err := s.DeleteSubs(tt.args.subs)
			assert.Equal(t, err != nil, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestSubsService_LoadSubsByID(t *testing.T) {
	subsModelFakeOk := FakeSubsModel{}
	subsModelFakeErr := FakeSubsErrorModel{}
	type fields struct {
		subsRepository repository.SubsRepository
	}
	type args struct {
		id primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Subs
		wantErr bool
	}{
		{"Ususal load subs", fields{&subsModelFakeOk}, args{testSubsID}, &models.Subs{}, false},
		{"Error load subs", fields{&subsModelFakeErr}, args{testSubsID}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubsServiceImpl{
				subsRepository: tt.fields.subsRepository,
			}
			got, err := s.LoadSubsByID(tt.args.id)
			assert.Equal(t, err != nil, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestSubsService_FindSubsByUID(t *testing.T) {
	subsModelFakeOk := FakeSubsModel{}
	subsModelFakeErr := FakeSubsErrorModel{}
	type fields struct {
		subsRepository repository.SubsRepository
	}
	type args struct {
		uid primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]models.Subs
		wantErr bool
	}{
		{"Ususal find by uid subs", fields{&subsModelFakeOk}, args{testSubsUID}, &[]models.Subs{}, false},
		{"Error find by uid subs", fields{&subsModelFakeErr}, args{testSubsUID}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubsServiceImpl{
				subsRepository: tt.fields.subsRepository,
			}
			got, err := s.FindSubsByUID(tt.args.uid)
			assert.Equal(t, err != nil, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestSubsService_FindSubsBetweenTime(t *testing.T) {
	subsModelFakeOk := FakeSubsModel{}
	subsModelFakeErr := FakeSubsErrorModel{}
	type fields struct {
		subsRepository repository.SubsRepository
		lastTime       int64
	}
	type args struct {
		fromTime int64
		toTime   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]models.Subs
		wantErr bool
	}{
		{"Ususal find by uid subs", fields{&subsModelFakeOk, 0}, args{0, 100000000}, &[]models.Subs{}, false},
		{"Error find by uid subs", fields{&subsModelFakeErr, 0}, args{0, 100000000}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubsServiceImpl{
				subsRepository: tt.fields.subsRepository,
				lastTime:       tt.fields.lastTime,
			}
			got, err := s.FindSubsBetweenTime(tt.args.fromTime, tt.args.toTime)
			assert.Equal(t, err != nil, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}
