package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getByIDCmd = &cobra.Command{
	Use:   "getbyid",
	Short: "Retourne un contact par son ID",
	Long:  `Retourne un contact spécifique en utilisant son ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil || id <= 0 {
			cmd.Help()
			return
		}
		contact, err := store.GetByID(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if contact == nil {
			fmt.Printf("Contact introuvable.\n")
			return
		}

		fmt.Printf("ID: %d, Name: %s, Email: %s\n", contact.ID, contact.Name, contact.Email)

	},
}
