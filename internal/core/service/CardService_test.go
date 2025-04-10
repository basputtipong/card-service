package service_test

import (
	"card-service/internal/core/domain"
	"card-service/internal/core/port"
	"card-service/internal/core/port/mocks"
	"card-service/internal/core/service"
	"testing"

	liberror "github.com/basputtipong/library/error"
	"github.com/stretchr/testify/assert"
)

func TestCardSvc_Execute(t *testing.T) {
	t.Run("Should_Pass_Return_Cards", func(t *testing.T) {
		mockRepo := new(mocks.CardRepo)
		svc := service.NewCardService(mockRepo)
		mockRepo.On("GetByUserId", "user-test").Return([]port.CardRepoRes{
			{
				CardID:      "1",
				UserID:      "user-test",
				Name:        "card1",
				Issuer:      "iss1",
				Number:      "1234",
				Status:      "Active",
				Color:       "#ffffff",
				BorderColor: "#ffffff",
			},
			{
				CardID:      "2",
				UserID:      "user-test",
				Name:        "card2",
				Issuer:      "iss2",
				Number:      "5678",
				Status:      "Inactive",
				Color:       "#ffffff",
				BorderColor: "#ffffff",
			},
		}, nil)

		expected := domain.CardRes{
			Cards: []domain.Card{
				{
					CardID:      "1",
					UserID:      "user-test",
					Name:        "card1",
					Issuer:      "iss1",
					Number:      "1234",
					Status:      "Active",
					Color:       "#ffffff",
					BorderColor: "#ffffff",
				},
				{
					CardID:      "2",
					UserID:      "user-test",
					Name:        "card2",
					Issuer:      "iss2",
					Number:      "5678",
					Status:      "Inactive",
					Color:       "#ffffff",
					BorderColor: "#ffffff",
				},
			},
		}
		req := domain.CardReq{
			UserId: "user-test",
		}
		res, err := svc.Execute(req)

		assert.NoError(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Should_Error", func(t *testing.T) {
		mockRepo := new(mocks.CardRepo)
		svc := service.NewCardService(mockRepo)
		mockRepo.On("GetByUserId", "user-test").Return([]port.CardRepoRes{}, liberror.ErrorInternalServerError("", ""))

		req := domain.CardReq{
			UserId: "user-test",
		}
		res, err := svc.Execute(req)

		expected := domain.CardRes{}
		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Should_Error_When_No_User_ID", func(t *testing.T) {
		mockRepo := new(mocks.CardRepo)
		svc := service.NewCardService(mockRepo)
		req := domain.CardReq{
			UserId: "",
		}
		res, err := svc.Execute(req)

		expected := domain.CardRes{}
		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})
}
