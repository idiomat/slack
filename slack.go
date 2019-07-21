package slack

const (
	// ResponseTypeInChannel is the in_channel response.
	ResponseTypeInChannel = "in_channel"

	// ResponseTypeEphemeral is the ephemeral response visible only to the interacting user.
	ResponseTypeEphemeral = "ephemeral"
)

const (
	// ColorBley is "bley"
	ColorBley = "#6B758B"
	// ColorYorange is "yorange"
	ColorYorange = "#FFB11D"
	// ColorInfo is an arbitrary blue color
	ColorInfo = "#00ADEE"
	// ColorBlack is black
	ColorBlack = "#000000"
	// ColorWarning is Slack's "warning" color
	ColorWarning = "warning"
	// ColorGood is Slack's "good" color
	ColorGood = "good"
	// ColorDanger is Slack's "danger" color
	ColorDanger = "danger"
)

// Field is part of the Attachment.
type Field struct {
	Title    string `json:"title"`
	Value    string `json:"value"`
	Short    bool   `json:"short"`
	Markdown bool   `json:"mrkdwn"`
}

// Action is part of the Attachment.
type Action struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Attachment is a part of a rich response to a Slack message.
type Attachment struct {
	Title          string    `json:"title"`
	TitleLink      string    `json:"title_link"`
	Fields         []*Field  `json:"fields,omitempty"`
	AuthorName     string    `json:"author_name,omitempty"`
	AuthorIcon     string    `json:"author_icon,omitempty"`
	ImageURL       string    `json:"image_url,omitempty"`
	ThumbURL       string    `json:"thumb_url,omitempty"`
	Text           string    `json:"text,omitempty"`
	Fallback       string    `json:"fallback,omitempty"`
	CallbackID     string    `json:"callback_id,omitempty"`
	Color          string    `json:"color,omitempty"`
	AttachmentType string    `json:"attachment_type,omitempty"`
	Footer         string    `json:"footer"`
	FooterIcon     string    `json:"footer_icon"`
	Actions        []*Action `json:"actions,omitempty"`
}

// Response is the response sent to Slack after capturing the request.
type Response struct {
	ResponseType string        `json:"response_type"`
	Text         string        `json:"text"`
	Attachments  []*Attachment `json:"attachments"`
}
