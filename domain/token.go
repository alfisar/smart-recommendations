package domain

type TokenLogin struct {
	ID      int    `json:"id"`
	Token   string `json:"token"`
	Expired string `json:"exp"`
}
