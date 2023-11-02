/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethanmick/lime/db"
	"github.com/ethanmick/lime/email"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var localFlag bool

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clean called")
		clean()
	},
}

func parse(byt []byte) error {
	var e email.Email
	err := json.Unmarshal(byt, &e)
	if err != nil {
		return err
	}

	fmt.Printf("%s: '%s'\n", e.FromEmail(), e.Subject)
	fmt.Printf("==> Is Newsletter: %v\n", email.IsNewsletter(&e))
	fmt.Printf("==> Is List: %v\n", email.IsList(&e))
	if e.FromEmail() == "hello@phind.com" {
		spew.Dump(e)
		decoded, _ := base64.URLEncoding.DecodeString(e.Body)
		spew.Dump(string(decoded[:]))
	}
	fmt.Println("")

	return nil
}

func clean() {
	db, err := db.GetDB()
	if err != nil {
		log.Fatalf("unable to retrieve database: %v", err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("emails"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := parse(v); err != nil {
				log.Fatalf("unable to parse email: %v", err)
			}
		}
		return nil
	})

}

func init() {
	cleanCmd.Flags().BoolVarP(&localFlag, "local", "l", false, "Run the command in local mode")
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
