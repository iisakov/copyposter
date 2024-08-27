package requester

type Response struct {
	Count int   `json:"count"`
	Items Items `json:"items"`
}

type Body struct {
	Response Response `json:"response"`
}
