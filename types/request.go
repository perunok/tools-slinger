package types

type Request struct {
	Body    string     `json:"body"`
	Headers []KeyValue `json:"headers"`
	URL     string     `json:"url"`
}
