package email

func IsList(e *Email) bool {
	return e.Headers["List-Id"] != ""
}
