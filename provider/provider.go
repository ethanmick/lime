package providers

import (
	"context"

	"github.com/ethanmick/lime/email"
)

type GetEmailsRequest struct {
	Limit     int
	PageToken string
}

type GetEmailsResponse struct {
	Emails    []email.Email
	NextToken string
	Total     int64
}

type Provider interface {
	GetEmails(context.Context, *GetEmailsRequest) (*GetEmailsResponse, error)
}
