package email

import "google.golang.org/api/gmail/v1"

type Headers map[string]string

type Email struct {
	// Id: The immutable ID of the message.
	Id string `json:"id,omitempty"`

	// Headers: List of headers on this email
	Headers Headers `json:"headers,omitempty"`

	// Subject: The subject of the message.
	Subject string `json:"subject,omitempty"`

	// From: List of addresses from the `From` header.
	From []string `json:"from,omitempty"`

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

func FromGmail(mes *gmail.Message) *Email {
	e := &Email{}
	e.Id = mes.Id
	e.Headers = make(Headers)
	for _, h := range mes.Payload.Headers {
		e.Headers[h.Name] = h.Value
	}
	e.Subject = e.Headers["Subject"]
	e.To = []string{e.Headers["To"]}
	e.From = []string{e.Headers["From"]}
	e.CC = []string{e.Headers["CC"]}
	e.BCC = []string{e.Headers["BCC"]}
	var body string
	for _, p := range mes.Payload.Parts {
		if p.MimeType == "text/html" {
			body = p.Body.Data
		}
	}
	e.Body = body
	e.Filename = mes.Payload.Filename
	e.Snippet = mes.Snippet
	return e
}
