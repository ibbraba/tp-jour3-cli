package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Met à jour un contact existant",
	Long:  `Met à jour les informations d'un contact existant dans le CRM.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Verification des flags
		id, err := cmd.Flags().GetInt("id")
		if err != nil || id <= 0 {
			cmd.Help()
			return
		}

		name := cmd.Flag("name").Value.String()
		email := cmd.Flag("email").Value.String()

		if name == "" && email == "" {
			cmd.Help()
			return
		}

		newName := cmd.Flag("name").Value.String()
		newEmail := cmd.Flag("email").Value.String()

		//Modifie le contact
		err = store.Update(id, newName, newEmail)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Contact avec l'ID %d mis à jour avec succès.\n", id)
	},
}
