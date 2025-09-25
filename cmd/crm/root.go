package cmd

import (
	"fmt"
	"os"

	"github.com/ibbraba/tp-jour3-cli/internal/storage"
	"github.com/spf13/cobra"
)

// Commande racine
var rootCmd = &cobra.Command{
	Use:   "contact-manager",
	Short: "Commande racine pour gérer les contacts",
	Long: `Commande racine pour gérer les contacts. vous pouvez utiliser les sous-commandes
pour ajouter, lister, mettre à jour ou supprimer des contacts.`,
}

var store storage.Storer = storage.NewGormStore()

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

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

	// Initialize your command here
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(getallCmd)
	rootCmd.AddCommand(getByIDCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)

	// Flags add command
	addCmd.Flags().StringP("name", "n", "", "Name of the contact")
	addCmd.Flags().StringP("email", "e", "", "Email of the contact")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")

	// Flag getByID command
	getByIDCmd.Flags().IntP("id", "i", 0, "ID of the contact")
	getByIDCmd.MarkFlagRequired("id")

	// Flags update command
	updateCmd.Flags().IntP("id", "i", 0, "ID of the contact")
	updateCmd.Flags().StringP("name", "n", "", "Name of the contact")
	updateCmd.Flags().StringP("email", "e", "", "Email of the contact")
	updateCmd.MarkFlagRequired("id")

	// Flag delete command
	deleteCmd.Flags().IntP("id", "i", 0, "ID of the contact")
	deleteCmd.MarkFlagRequired("id")

}
