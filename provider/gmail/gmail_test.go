package gmail_test

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"

	gprovider "github.com/ethanmick/lime/provider/gmail"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/gmail/v1"
)

// write a test for the FromGmail method

func TestFromGmail(t *testing.T) {
	var testCases = []struct {
		name string
		file string
	}{
		{"simple email", "email-0.json"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonFile, err := os.Open(filepath.Join("testdata", tc.file))
			if err != nil {
				t.Fatal(err)
			}
			defer jsonFile.Close()

			byteValue, err := io.ReadAll(jsonFile)
			if err != nil {
				t.Fatal(err)
			}

			var message gmail.Message
			err = json.Unmarshal(byteValue, &message)
			if err != nil {
				t.Fatal(err)
			}

			mes := gprovider.FromGmail(&message)
			assert.NotEmpty(t, mes.Body, "The body should not be empty")
		})
	}
}
