package entities

type Evaluation struct {
	ID               uint `json:"id" gorm:"primaryKey"`
	EventDetailsID   EventDetails
	EventDetails     EventDetails
	Product          string `json:"product"`
	ProductIndicator string `json:"product_indicator"`
	ProductTarget    string `json:"product_target"`
	Result           string `json:"result"`
	ResultIndicator  string `json:"result_indicator"`
	ResultTarget     string `json:"result_target"`
}
