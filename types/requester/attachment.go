package requester

type Attachment struct {
	Type string `json:"type"`
	AttachmentInterface
}

type Attachments []Attachment

type AttachmentInterface interface{}
