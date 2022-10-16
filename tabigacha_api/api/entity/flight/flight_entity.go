package entity

type Section struct {
	Origin      string
	Destination string
}

type FlightEntity struct {
	BestPrice     map[string]string `dynamodbav:"BestPrice"`
	PriceCalender map[string]string `dynamodbav:"PriceCalender"`
}

type BestPriceEntity struct {
	Price int
	URL   string
}

type PriceCalenderEntity struct {
	PriceCalender []PriceCalender
}

type PriceCalender struct {
	Date  string
	Price int
}
