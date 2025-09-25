package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprime un contact",
	Long:  `Supprime un contact du CRM en utilisant son ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Verification du flag
		id, err := cmd.Flags().GetInt("id")
		if err != nil || id <= 0 {
			cmd.Help()
			return
		}

		// Supprime le contact
		err = store.Delete(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Contact avec l'ID %d supprimé avec succès.\n", id)
	},
}
