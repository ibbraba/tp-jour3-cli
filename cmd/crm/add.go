package cmd

import (
	"fmt"

	"github.com/ibbraba/tp-jour3-cli/internal/storage"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajoute un nouveau contact",
	Long:  `Ajoute un nouveau contact au CRM.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		email := cmd.Flag("email").Value.String()
		// Verification des flags
		if name == "" || email == "" {
			cmd.Help()
			return
		}

		// Transforme les flags en contact
		contact := &storage.Contact{
			Name:  name,
			Email: email,
		}

		// Ajoute le contact au storer
		err := store.Add(contact)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Contact '%s' ajouté avec l'ID %d.\n", contact.Name, contact.ID)

	},
}
