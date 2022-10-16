package response

type PriceCalenderResponse struct {
	Dates []DatePrice `json:"dates"`
}

type DatePrice struct {
	Price int    `json:"price"`
	Date  string `json:"date"`
}
