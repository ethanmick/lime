package stats

import "github.com/ethanmick/lime/email"

type DomainStat struct {
	Domain   string `json:"domain,omitempty"`
	Received int    `json:"received"`
	Send     int    `json:"sent"`

	Accounts map[string]SenderStat
}

type SenderStat struct {
	Sender   string `json:"domain,omitempty"`
	Received int    `json:"received"`
}

type Stats struct {
	// Total number of emails received
	Total int64
	// DomainStats for all domains
	DomainStats map[string]DomainStat
}

// Calculate the stats for all passed emails
func Calculate(emails []email.Email) *Stats {

	// For each email

	// Add the stats for the domain

	return nil

}
