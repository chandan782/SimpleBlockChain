package transaction

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Transaction struct {
	ID         string `json:"buyer" gorm:"column:"`
	BuyerID    string `json:"buyerId" gorm:"column:buyerId"`
	SellerID   string `json:"sellerId" gorm:"column:sellerId"`
	Status     string `json:"status" gorm:"column:status"`
	Securities string `json:"securities" gorm:"column:securities"`
	CreatedAt  string `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt  string `json:"updatedAt" gorm:"column:updatedAt"`
	Method     string `json:"method" gorm:"column:method"`
}

type Security struct {
	ID       string `json:"id" gorm:"column:id"`
	Status   string `json:"status" gorm:"column:status"`
	ScriptID string `json:"scriptId" gorm:"column:scriptId"`
	Quantity int    `json:"quantity" gorm:"column:quantity"`
	Amount   int    `json:"amount" gorm:"column:amount"`
}
