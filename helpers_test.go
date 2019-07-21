package slack_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/idiomat/slack"
)

func TestGenSlackErrorResponse(t *testing.T) {
	tests := map[string]struct {
		input       error
		resp        *slack.Response
		errExpected error
	}{
		"success": {
			input: errors.New("An error occured"),
			resp: &slack.Response{
				ResponseType: slack.ResponseTypeEphemeral,
				Attachments: []*slack.Attachment{
					&slack.Attachment{
						Title: "Error",
						Text:  "An error occured",
						Color: slack.ColorDanger,
					},
				},
			},
		},
		"failure": {
			input:       nil,
			resp:        nil,
			errExpected: errors.New("err cannot be nil"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, err := slack.GenSlackErrorResponse(tc.input)

			if tc.errExpected != nil {
				if diff := cmp.Diff(tc.errExpected.Error(), err.Error()); diff != "" {
					t.Fatalf("Expected error (-want +got):\n%s", diff)
				}
			}

			if diff := cmp.Diff(tc.resp, r); diff != "" {
				t.Errorf("Expected a response (-want +got):\n%s", diff)
			}
		})
	}
}
