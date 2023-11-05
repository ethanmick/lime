/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"log"
	"log/slog"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethanmick/lime/classifier"
	"github.com/ethanmick/lime/db"
	"github.com/ethanmick/lime/email"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// labelCmd represents the label command
var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("wtf")
		label()
	},
}

func label() {
	slog.Debug("Here 0")
	classy := classifier.NewOpenAIClassifier()

	db, err := db.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("emails"))
		c := b.Cursor()
		count := 0
		for k, v := c.First(); k != nil; k, v = c.Next() {
			slog.Debug("Here 1")
			count++
			var e email.Email
			err := json.Unmarshal(v, &e)
			if err != nil {
				return err
			}

			labels, err := classy.Classify(&e)
			if err != nil {
				return err
			}

			if len(labels) == 0 {
				slog.Error("No labels found for email", "subject", e.Subject)
				spew.Dump(e)
			}

			e.Labels = append(e.Labels, labels...)

			toSave, err := json.Marshal(e)
			if err != nil {
				log.Fatalf("unable to marshal email: %v", err)
			}

			err = b.Put([]byte(e.ID), toSave)
			if err != nil {
				log.Fatalf("unable to save email: %v", err)
			}
			if count > 10 {
				break
			}
			slog.Debug("Sleeping 20 seconds...")
			time.Sleep(20 * time.Second)
		}
		return nil
	})

	if err != nil {
		slog.Error("Error", err)
	}

}

func init() {
	rootCmd.AddCommand(labelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// labelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// labelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
