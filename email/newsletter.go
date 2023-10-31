package email

func IsNewsletter(e *Email) bool {
	return e.Headers["List-Unsubscribe"] != ""
}

func GetNewsletterName(e *Email) string {
	return ""
}
