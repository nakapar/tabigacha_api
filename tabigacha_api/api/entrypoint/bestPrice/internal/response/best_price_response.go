package response

type BestPriceResponse struct {
	Price int      `json:"price"`
	Date  []string `json:"date"`
}
