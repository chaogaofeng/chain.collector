package schema

type Token struct {
	ID          uint   `json:"-" gorm:"primarykey" `
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Decimal     int    `json:"decimal"`
	Display     string `json:"display"`
	denom       string `json:"denom"`
	TotalSupply string `json:"total_supply"`
	Supply      string `json:"supply"`
	Issued      string `json:"issued"`
	Burned      string `json:"burned"`
	Owner       string `json:"owner"`
}
