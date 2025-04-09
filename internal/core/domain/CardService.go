package domain

import "card-service/internal/core/port"

type CardService interface {
	Execute(req CardReq) (CardRes, error)
}

type CardReq struct {
	UserId string `json:"userId" validate:"required"`
}

type CardRes struct {
	CardID      string `json:"cardId"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	Issuer      string `json:"issuer"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	Color       string `json:"color"`
	BorderColor string `json:"borderColor"`
}

func (res *CardRes) BuildCardResponse(repoRes port.CardRepoRes) {
	res.CardID = repoRes.CardID
	res.UserID = repoRes.UserID
	res.Name = repoRes.Name
	res.Issuer = repoRes.Issuer
	res.Number = repoRes.Number
	res.Status = repoRes.Status
	res.Color = repoRes.Color
	res.BorderColor = repoRes.BorderColor
}
