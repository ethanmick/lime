/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	oauthv2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your email account",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		login()
	},
}

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	var wg sync.WaitGroup
	shutdown := make(chan struct{})
	server := &http.Server{Addr: "localhost:6925"}
	var authCode string

	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Page fetched: %s\n", r.URL)
		authCode = r.URL.Query().Get("code")

		w.Header().Set("Content-Type", "text/html")
		responseHTML := `<html><body><p>You can close this tab now.</p></body></html>`
		_, err := w.Write([]byte(responseHTML))
		if err != nil {
			log.Printf("Failed to send response: %v", err)
		}

		close(shutdown)
	})

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	go func() {
		// Wait for the shutdown signal.
		<-shutdown

		// We use a sync.WaitGroup to ensure the server has finished processing the request.
		wg.Add(1)

		// Shutdown the server gracefully.
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Could not shutdown: %v", err)
		}

		wg.Done()
	}()

	err := exec.Command("open", authURL).Run()
	if err != nil {
		log.Fatalf("Unable to open browser: %v", err)
	}

	// Start the server.
	wg.Add(1)
	go func() {
		fmt.Println("Server is listening on localhost:6925")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}

		wg.Done()
	}()

	// Wait for the server to finish.
	wg.Wait()

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// Retrieves a token from a local file.
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

func login() {
	ctx := context.Background()

	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.MailGoogleComScope, "email")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	tok, _ := tokenFromFile("token.json")
	spew.Dump(tok)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	email := data["email"].(string)
	viper.Set("account", email)
	viper.WriteConfig()

	spew.Dump(email, data)

	// client := conf.Client(context.Background(), tok)

	srv, err := oauthv2.NewService(ctx,
		option.WithTokenSource(config.TokenSource(ctx, tok)))
	if err != nil {
		log.Fatalf("Unable to retrieve OAuth2 client: %v", err)
	}

	userinfo, err := srv.Userinfo.Get().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve userinfo: %v", err)
	}

	spew.Dump(userinfo)

}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
