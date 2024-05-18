package models

type Transaction struct {
    ToAddress string `json:"to_address"`
    Amount    string `json:"amount"`
}

type Batch struct {
    ID           string        `json:"id"`
    Transactions []Transaction `json:"transactions"`
}
