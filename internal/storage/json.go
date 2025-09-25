package storage

import (
	"encoding/json"
	"os"
)

type JSONStore struct {
	filePath string
}

func NewJSONStore(filePath string) *JSONStore {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Crée le fichier
		f, err := os.Create(filePath)

		if err != nil {
			panic("Impossible de créer le fichier de sortie JSON: " + err.Error())
		}
		// Insérer un tableau vide dans le fichier
		_, err = f.WriteString("[]")
		if err != nil {
			panic("Impossible d'écrire dans le fichier de sortie JSON: " + err.Error())
		}
		f.Close()
	}

	return &JSONStore{
		filePath: filePath,
	}
}

func (js *JSONStore) Add(contact *Contact) error {

	//Lis le fichier contacts.json s'il existe
	var contacts []*Contact
	fileData, err := os.ReadFile(js.filePath)
	if err == nil && len(fileData) > 0 {
		json.Unmarshal(fileData, &contacts)
	}

	// Détermine le prochain ID
	var maxID int
	for _, c := range contacts {
		if c.ID > maxID {
			maxID = c.ID
		}
	}
	contact.ID = maxID + 1

	//Ajoute le contact au slice
	contacts = append(contacts, contact)

	// Met à jour JSON data
	jsonData, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return err
	}

	// Ecrit Json data à nouveau dans le fichier
	osWrite := os.WriteFile(js.filePath, jsonData, 0644)
	if osWrite != nil {
		return osWrite
	}
	return nil
}

func (js *JSONStore) GetAll() ([]*Contact, error) {
	fileData, err := os.ReadFile(js.filePath)
	if err != nil {
		return nil, err
	}

	var contacts []*Contact
	if err := json.Unmarshal(fileData, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}

func (js *JSONStore) GetByID(id int) (*Contact, error) {
	//Lis le fichier contacts.json s'il existe
	var contacts []*Contact
	fileData, err := os.ReadFile(js.filePath)
	if err == nil && len(fileData) > 0 {
		json.Unmarshal(fileData, &contacts)
	}

	if len(contacts) == 0 {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	//Cherche le contact avec l'ID donné

	for _, contact := range contacts {
		if contact.ID == id {
			return contact, nil
		}
	}
	return nil, nil
}

func (js *JSONStore) Update(id int, newName, newEmail string) error {
	contacts, err := js.GetAll()
	if err != nil {
		return err
	}

	for i, contact := range contacts {
		if contact.ID == id {
			contacts[i].Name = newName
			contacts[i].Email = newEmail
			break
		}
	}

	jsonData, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(js.filePath, jsonData, 0644); err != nil {
		return err
	}
	return nil
}

func (js *JSONStore) Delete(id int) error {
	contacts, err := js.GetAll()
	if err != nil {
		return err
	}

	var updatedContacts []*Contact
	for _, contact := range contacts {
		if contact.ID != id {
			updatedContacts = append(updatedContacts, contact)
		}
	}

	jsonData, err := json.MarshalIndent(updatedContacts, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(js.filePath, jsonData, 0644); err != nil {
		return err
	}
	return nil
}
