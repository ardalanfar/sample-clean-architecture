package dto

/*--------------------Show Orders--------------------*/

type ShowOrdersRequest struct {
	ID uint `json:"id"`
}

type ShowOrdersResponse struct {
	Orders []Showorders `json:"orders"`
}
