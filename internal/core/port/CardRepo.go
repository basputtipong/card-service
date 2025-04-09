package port

type CardRepo interface {
	GetByUserId(userId string) (CardRepoRes, error)
}

const (
	DebitCardsTbl       = "debit_cards"
	DebitCardDetailsTbl = "debit_card_details"
	DebitCardStatusTbl  = "debit_card_status"
	DebitCardDesignTbl  = "debit_card_design"
)

type CardRepoRes struct {
	CardID      string `gorm:"column:card_id"`
	UserID      string `gorm:"column:user_id"`
	Name        string `gorm:"column:name"`
	Issuer      string `gorm:"column:issuer"`
	Number      string `gorm:"column:number"`
	Status      string `gorm:"column:status"`
	Color       string `gorm:"column:color"`
	BorderColor string `gorm:"column:border_color"`
}
