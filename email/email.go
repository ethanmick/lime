package email

import (
	"net/mail"
	"time"
)

type Headers map[string]string

type Email struct {
	// Id: The immutable ID of the message.
	ID string `json:"id,omitempty"`

	// Received: The time at which the message was received by the server
	Received *time.Time `json:"received,omitempty"`

	// Headers: List of headers on this email
	Headers Headers `json:"headers,omitempty"`

	// Subject: The subject of the message.
	Subject string `json:"subject,omitempty"`

	// From: List of addresses from the `From` header.
	From string `json:"from,omitempty"`

	// To: List of addresses from the `To` header.
	To []string `json:"to,omitempty"`

	// CC: List of addresses from the `CC` header.
	CC []string `json:"cc,omitempty"`

	// BCC: List of addresses from the `BCC` header.
	BCC []string `json:"bcc,omitempty"`

	// Body: The entire email message in an RFC 2822 formatted and base64url
	Body string `json:"body,omitempty"`

	// Filename: The filename of the attachment. Only present if this
	// message part represents an attachment.
	Filename string `json:"filename,omitempty"`

	// Snippet: A short part of the message text.
	Snippet string `json:"snippet,omitempty"`
}

func (e *Email) FromEmail() string {
	addr, err := mail.ParseAddress(e.From)
	if err != nil {
		return ""
	}
	return addr.Address
}
