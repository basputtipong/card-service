package domain

import "card-service/internal/core/port"

type CardService interface {
	Execute(req CardReq) (CardRes, error)
}

type CardReq struct {
	UserId string `json:"userId" validate:"required"`
}

type CardRes struct {
	Cards []Card `json:"cards"`
}

type Card struct {
	CardID      string `json:"cardId"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	Issuer      string `json:"issuer"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	Color       string `json:"color"`
	BorderColor string `json:"borderColor"`
}

func (res *CardRes) BuildCardResponse(repoRes []port.CardRepoRes) {
	var cards []Card
	for _, ele := range repoRes {
		cards = append(cards, Card{
			CardID:      ele.CardID,
			UserID:      ele.UserID,
			Name:        ele.Name,
			Issuer:      ele.Issuer,
			Number:      ele.Number,
			Status:      ele.Status,
			Color:       ele.Color,
			BorderColor: ele.BorderColor,
		})
	}
	res.Cards = cards
}
