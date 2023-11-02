package gmail

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ethanmick/lime/email"
	providers "github.com/ethanmick/lime/provider"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailProvider struct {
	srv *gmail.Service
}

func FromGmail(mes *gmail.Message) email.Email {
	e := email.Email{}
	e.ID = mes.Id
	e.Headers = make(email.Headers)
	t := time.Unix(mes.InternalDate/1000, 0)
	e.Received = &t

	for _, h := range mes.Payload.Headers {
		e.Headers[h.Name] = h.Value
	}
	e.Subject = e.Headers["Subject"]
	e.To = []string{e.Headers["To"]}
	e.From = e.Headers["From"]
	e.CC = []string{e.Headers["Cc"]}
	e.BCC = []string{e.Headers["Bcc"]}
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

func (g *GmailProvider) GetEmails(c context.Context, req *providers.GetEmailsRequest) (*providers.GetEmailsResponse, error) {
	user := "me"
	pageToken := ""
	response := &providers.GetEmailsResponse{}
	count := 0
	for {
		gmailReq := g.srv.Users.Messages.List(user).PageToken(pageToken)
		if req.Limit > 0 {
			gmailReq = gmailReq.MaxResults(int64(req.Limit))
		}
		resp, err := gmailReq.Do()
		if err != nil {
			return nil, fmt.Errorf("failed to list user messages: %w", err)
		}
		response.Total = resp.ResultSizeEstimate
		for _, m := range resp.Messages {
			slog.Info(fmt.Sprintf("Fetching message %d/%d", count, response.Total))
			msg, err := g.srv.Users.Messages.Get("me", m.Id).Format("full").Do()
			if err != nil {
				return nil, fmt.Errorf("failed to get message: %w", err)
			}
			response.Emails = append(response.Emails, FromGmail(msg))
			count++
		}

		if len(resp.Messages) == 0 || len(response.Emails) >= req.Limit {
			break
		}

		pageToken = resp.NextPageToken
	}

	return response, nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func getClient(config *oauth2.Config) (*http.Client, error) {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	return config.Client(context.Background(), tok), err
}

func NewClient(c context.Context) (*gmail.Service, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// 	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client, err := getClient(config)
	if err != nil {
		log.Fatalf("Login token not found or invalid, please run `lime login` to generate a new token: %v", err)
	}

	srv, err := gmail.NewService(c, option.WithHTTPClient(client))
	return srv, err
}

func NewProvider(c context.Context) (providers.Provider, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}
	client, err := getClient(config)
	if err != nil {
		return nil, fmt.Errorf("login token not found or invalid, please run `lime login` to generate a new token: %w", err)
	}
	srv, err := gmail.NewService(c, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Gmail client: %w", err)
	}

	return &GmailProvider{srv}, nil

}
