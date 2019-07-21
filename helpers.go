package slack

import "errors"

// GenSlackErrorResponse generates an error payload for Slack to render on the client.
func GenSlackErrorResponse(err error) (*Response, error) {
	if err == nil {
		return nil, errors.New("err cannot be nil")
	}

	return &Response{
		ResponseType: ResponseTypeEphemeral,
		Attachments: []*Attachment{
			&Attachment{
				Title: "Error",
				Text:  err.Error(),
				Color: ColorDanger,
			},
		},
	}, nil
}
