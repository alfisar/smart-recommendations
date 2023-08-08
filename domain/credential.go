package domain

type Credential struct {
	ID          int    `json:"id" grom:"column:id"`
	Application string `json:"application" grom:"column:application"`
	Password    string `json:"password" grom:"column:password"`
	Description string `json:"description" grom:"column:description"`
}
