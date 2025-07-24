package types

type Response struct {
	Body       string     `json:"body"`
	Headers    []KeyValue `json:"headers"`
	StatusCode int        `json:"status-code"`
}
