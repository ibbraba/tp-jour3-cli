package cmd

import (
	"fmt"
	"os"

	"github.com/ibbraba/tp-jour3-cli/internal/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Commande racine
var rootCmd = &cobra.Command{
	Use:   "contact-manager",
	Short: "Commande racine pour gérer les contacts",
	Long: `Commande racine pour gérer les contacts. vous pouvez utiliser les sous-commandes
pour ajouter, lister, mettre à jour ou supprimer des contacts.`,
}

var store storage.Storer = storage.NewGormStore()

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)
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

func initConfig() {

	// Configuration de Viper pour lire le fichier config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error reading config file:", err)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Lis le fichier config.yaml pour savoir quel type de stockage utiliser
	driver := viper.GetString("storage.driver")
	switch driver {
	case "gorm":
		store = storage.NewGormStore()
	case "json":
		store = storage.NewJSONStore("contacts.json")
	case "memory":
		store = storage.NewMemoryStore()
	default:
		panic("Veuillez spécifier un driver de stockage valide dans le fichier config.yaml")
	}
}
