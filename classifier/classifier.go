package classifier

import "github.com/ethanmick/lime/email"

type Classifer interface {
	Classify(*email.Email) ([]email.Label, error)
}
