package types

// Money Денежная сумма
type Money int64

// Currency валюта
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN номер карты
type PAN string

// Card платежная карта
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Category категория
type Category string

// Status статус платежа
type Status string

// статусы платежа
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment информация о платеже
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
