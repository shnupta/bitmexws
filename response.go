package bitmexws

// Represents a success message sent back by BitMEX after subscribing to a
// topic.
// Example:
//	{"success":true,"subscribe":"trade:XBTUSD",
//     "request":{"op":"subscribe","args":["trade:XBTUSD","instrument:XBTUSD"]}}
type SubSuccess struct {
	Subscribe string `json:"subscribe"`
	Success   bool   `json:"success"`
}

// Represents an error message sent back by BitMEX after subscribing to a
// topic.
// Example:
//   {"error": errorMessage}
type SubError struct {
	Error string `json:"error"`
}

type SubFilter struct {
	Account int    `json:"account"`
	Symbol  string `json:"symbol"`
}

// Represents a data message sent back by BitMEX after subscribing to a
// topic.
// Example:
//   {"table":"instrument","action":"update","data":[{"symbol":"XBTUSD","openValue":673520400,
//     "timestamp":"2016-01-15T22:33:15.000Z"}]}
type SubData struct {
	Table       string            `json:"table"`
	Action      string            `json:"action"`
	Data        []interface{}     `json:"data"`
	Keys        []string          `json:"keys"`
	ForeignKeys map[string]string `json:"foreignKeys"`
	Types       map[string]string `json:"types"`
	Filter      SubFilter         `json:"filter"`
	Attributes  map[string]string `json:"attributes"`
}
