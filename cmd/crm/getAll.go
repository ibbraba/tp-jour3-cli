package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getallCmd = &cobra.Command{
	Use:   "getall",
	Short: "Retourne tous les contacts",
	Long:  `Retourne tous les contacts du CRM.`,
	Run: func(cmd *cobra.Command, args []string) {
		contacts, err := store.GetAll()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		if len(contacts) == 0 {
			fmt.Println(" Pas de contacts à afficher.")
			return
		}

		fmt.Println("\n--- Liste des contacts ---")
		for _, contact := range contacts {
			fmt.Printf("ID: %d, Name: %s, Email: %s\n", contact.ID, contact.Name, contact.Email)
		}
	},
}
