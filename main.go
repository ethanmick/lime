package main

import "github.com/ethanmick/lime/cmd"

func main() {
	cmd.Execute()
}

// // Retrieve a token, saves the token, then returns the generated client.
// func getClient(config *oauth2.Config) *http.Client {
// 	// The file token.json stores the user's access and refresh tokens, and is
// 	// created automatically when the authorization flow completes for the first
// 	// time.
// 	tokFile := "token.json"
// 	tok, err := tokenFromFile(tokFile)
// 	if err != nil {
// 		tok = getTokenFromWeb(config)
// 		saveToken(tokFile, tok)
// 	}
// 	return config.Client(context.Background(), tok)
// }

// // Request a token from the web, then returns the retrieved token.
// func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	fmt.Printf("Go to the following link in your browser then type the "+
// 		"authorization code: \n%v\n", authURL)

// 	var authCode string
// 	if _, err := fmt.Scan(&authCode); err != nil {
// 		log.Fatalf("Unable to read authorization code: %v", err)
// 	}

// 	tok, err := config.Exchange(context.TODO(), authCode)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve token from web: %v", err)
// 	}
// 	return tok
// }

// // Retrieves a token from a local file.
// func tokenFromFile(file string) (*oauth2.Token, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	tok := &oauth2.Token{}
// 	err = json.NewDecoder(f).Decode(tok)
// 	return tok, err
// }

// // Saves a token to a file path.
// func saveToken(path string, token *oauth2.Token) {
// 	fmt.Printf("Saving credential file to: %s\n", path)
// 	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
// 	if err != nil {
// 		log.Fatalf("Unable to cache oauth token: %v", err)
// 	}
// 	defer f.Close()
// 	json.NewEncoder(f).Encode(token)
// }

// func main() {
// 	ctx := context.Background()
// 	b, err := os.ReadFile("credentials.json")
// 	if err != nil {
// 		log.Fatalf("Unable to read client secret file: %v", err)
// 	}

// 	// If modifying these scopes, delete your previously saved token.json.
// 	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
// 	if err != nil {
// 		log.Fatalf("Unable to parse client secret file to config: %v", err)
// 	}
// 	client := getClient(config)

// 	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve Gmail client: %v", err)
// 	}

// 	db, err := bolt.Open("my.db", 0600, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	db.Update(func(tx *bolt.Tx) error {
// 		_, err := tx.CreateBucketIfNotExists([]byte("emails"))
// 		if err != nil {
// 			return fmt.Errorf("create bucket: %s", err)
// 		}
// 		return nil
// 	})

// 	user := "me"
// 	r, err := srv.Users.Labels.List(user).Do()
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve labels: %v", err)
// 	}
// 	if len(r.Labels) == 0 {
// 		fmt.Println("No labels found.")
// 		return
// 	}
// 	fmt.Println("Labels:")
// 	for _, l := range r.Labels {
// 		fmt.Printf("- %s\n", l.Name)
// 	}

// 	var messages []*gmail.Message
// 	mes, err := srv.Users.Messages.List(user).MaxResults(500).Do()
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve labels: %v", err)
// 	}

// 	fmt.Println("Length", len(mes.Messages))
// 	for _, m := range mes.Messages {
// 		// fmt.Println(m.Id)
// 		messages = append(messages, m)
// 	}

// 	for _, m := range messages {
// 		msg, err := srv.Users.Messages.Get("me", m.Id).Format("full").Do()

// 		if err != nil {
// 			log.Fatalf("Unable to retrieve labels: %v", err)
// 		}
// 		var subject string
// 		for _, h := range msg.Payload.Headers {
// 			if h.Name == "Subject" {
// 				subject = h.Value
// 			}
// 		}

// 		fmt.Println("Subject:", subject)
// 		// base 64 decode the body
// 		// spew.Dump(msg.Payload.Parts)

// 		var data string
// 		for _, p := range msg.Payload.Parts {
// 			if p.MimeType == "text/html" {
// 				data = p.Body.Data
// 			}
// 		}
// 		// fmt.Println("wtf", data)
// 		decoded, err := base64.URLEncoding.DecodeString(data)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// fmt.Println(string(decoded[:]))

// 		// new reader from bytes
// 		r := io.Reader(bytes.NewReader(decoded))
// 		doc, err := goquery.NewDocumentFromReader(r)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		title := doc.Find("title").First()

// 		fmt.Println("Title", title.Text())
// 		return
// 	}

// }
