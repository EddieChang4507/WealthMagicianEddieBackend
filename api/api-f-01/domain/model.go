package domain

type Date struct {
	date string
}

// 開高低收
type CandlestickChart struct {
	date         string
	OpenPrice    int64 `json:"openprice"`
	HighestPrice int64 `json:"highestprice"`
	ClosingPrice int64 `json:"closingprice"`
	LowestPrice  int64 `json:"lowestprice"`
}
