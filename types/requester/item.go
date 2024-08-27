package requester

type Item struct {
	InnerType   string      `json:"inner_type"`
	Type        string      `json:"type"`
	Attachments Attachments `json:"attachments"`
	Date        int         `json:"date"`
	FromId      int         `json:"from_id"`
	Id          int         `json:"id"`
	OwnerId     int         `json:"owner_id"`
	Text        string      `json:"text"`
}

type Items []Item
