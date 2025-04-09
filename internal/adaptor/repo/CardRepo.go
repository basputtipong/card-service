package repository

import (
	"card-service/internal/core/port"
	"fmt"

	"gorm.io/gorm"
)

type cardRepo struct {
	db *gorm.DB
}

func NewCardRepo(db *gorm.DB) port.CardRepo {
	return &cardRepo{db}
}

func (r *cardRepo) GetByUserId(userId string) (port.CardRepoRes, error) {
	var repoRes port.CardRepoRes
	err := r.db.Table(port.DebitCardsTbl+" AS dc").
		Select("dc.card_id, dc.user_id, dc.name, dcd.issuer, dcd.number, dcs.status, dcn.color, dcn.border_color").
		Joins(fmt.Sprintf(`LEFT JOIN %s dcd ON dc.card_id = dcd.card_id`, port.DebitCardDetailsTbl)).
		Joins(fmt.Sprintf(`LEFT JOIN %s dcs ON dc.card_id = dcs.card_id`, port.DebitCardStatusTbl)).
		Joins(fmt.Sprintf(`LEFT JOIN %s dcn ON dc.card_id = dcn.card_id`, port.DebitCardDesignTbl)).
		Where("dc.user_id = ?", userId).
		Scan(&repoRes).Error
	if err != nil {
		return repoRes, err
	}
	return repoRes, nil
}
