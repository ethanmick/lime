package cmd

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ethanmick/lime/db"
	providers "github.com/ethanmick/lime/provider"
	"github.com/ethanmick/lime/provider/gmail"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var limit int

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		download()
	},
}

func download() {
	ctx := context.Background()
	client, err := gmail.NewProvider(ctx)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	opts := &providers.GetEmailsRequest{}
	if limit > 0 {
		opts.Limit = limit
	}

	resp, err := client.GetEmails(ctx, opts)
	if err != nil {
		log.Fatalf("unable to retrieve emails: %v", err)
	}

	db, err := db.GetDB()
	if err != nil {
		log.Fatalf("unable to retrieve database: %v", err)
	}
	defer db.Close()

	for _, email := range resp.Emails {
		log.Printf("Saving email: %s", email.ID)
		e, err := json.Marshal(email)
		if err != nil {
			log.Fatalf("unable to marshal email: %v", err)
		}
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("emails"))
			err := b.Put([]byte(email.ID), e)
			return err
		})
	}
}

func init() {
	downloadCmd.Flags().IntVarP(&limit, "limit", "", -1, "Limit the number of messages downloaded")
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
