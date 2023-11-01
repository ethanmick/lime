package email

type DomainStats struct {
	Domain        string `json:"domain"`
	FromCount     int64  `json:"from_count"`
	ResponseCount int64  `json:"response_count"`
}

type Statistics struct {
	TotalEmails int64                  `json:"total_emails"`
	DomainStats map[string]DomainStats `json:"domain_stats"`
}
