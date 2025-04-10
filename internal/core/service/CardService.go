package service

import (
	"card-service/internal/core/domain"
	"card-service/internal/core/port"
	"card-service/utils"

	liberror "github.com/basputtipong/library/error"
)

type cardSvc struct {
	cardRepo port.CardRepo
}

func NewCardService(cardRepo port.CardRepo) domain.CardService {
	return &cardSvc{cardRepo}
}

func (s *cardSvc) Execute(req domain.CardReq) (domain.CardRes, error) {
	var res domain.CardRes
	if err := utils.Validate(req); err != nil {
		return res, liberror.ErrorBadRequest("Invalid request", err.Error())
	}

	repoRes, err := s.cardRepo.GetByUserId(req.UserId)
	if err != nil {
		return res, liberror.ErrorInternalServerError("Failed when trying to get card", err.Error())
	}

	res.BuildCardResponse(repoRes)
	return res, nil
}
