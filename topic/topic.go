package topic

type Topics struct {
	Announcement         string
	Chat                 string
	Connected            string
	Funding              string
	Instrument           string
	Insurance            string
	Liquidation          string
	OrderBookL2_25       string
	OrderBookL2          string
	OrderBook10          string
	PublicNotifications  string
	Quote                string
	QuoteBin1m           string
	QuoteBin5m           string
	QuoteBin1h           string
	QuoteBin1d           string
	Settlement           string
	Trade                string
	TradeBin1m           string
	TradeBin5m           string
	TradeBin1h           string
	TradeBin1d           string
	Affiliate            string
	Execution            string
	Order                string
	Margin               string
	Position             string
	PrivateNotifications string
	Transact             string
	Wallet               string
}

// All possible op strings to send
func GetTopics() *Topics {
	// Authentication not required
	return &Topics{
		Announcement:         "announcement",
		Chat:                 "chat",
		Connected:            "connected",
		Funding:              "funding",
		Instrument:           "instrument",
		Insurance:            "insurance",
		Liquidation:          "liquidation",
		OrderBookL2_25:       "orderBookL2_25",
		OrderBookL2:          "orderBookL2",
		OrderBook10:          "orderBook10",
		PublicNotifications:  "publicNotifications",
		Quote:                "quote",
		QuoteBin1m:           "quoteBin1m",
		QuoteBin5m:           "quoteBin5m",
		QuoteBin1h:           "quoteBin1h",
		QuoteBin1d:           "quoteBin1d",
		Settlement:           "settlement",
		Trade:                "trade",
		TradeBin1m:           "tradeBin1m",
		TradeBin5m:           "tradeBin5m",
		TradeBin1h:           "tradeBin1h",
		TradeBin1d:           "tradeBin1d",
		Affiliate:            "affiliate",
		Execution:            "execution",
		Order:                "order",
		Margin:               "margin",
		Position:             "position",
		PrivateNotifications: "privateNotifications",
		Transact:             "transact",
		Wallet:               "wallet",
	}
}
