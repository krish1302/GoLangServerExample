package api

type Item struct{
	ID string `json:"id"`
	Product string `json:"product"`
	Price string `json:"price"`
}

type Login struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}