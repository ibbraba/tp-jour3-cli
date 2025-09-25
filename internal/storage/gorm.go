package storage

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore() *GormStore {
	db, err := gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		//Panic si DB inaccessible
		fmt.Println("failed to connect database")
		panic(err)
	}

	// Migration auto de la table Contact
	err = db.AutoMigrate(&Contact{})
	if err != nil {
		panic("failed to migrate: " + err.Error())
	}
	return &GormStore{db: db}
}

func (gs *GormStore) Add(contact *Contact) error {

	return gs.db.Create(contact).Error
}

func (gs *GormStore) GetAll() ([]*Contact, error) {
	var contacts []*Contact
	err := gs.db.Find(&contacts).Error
	return contacts, err
}

func (gs *GormStore) GetByID(id int) (*Contact, error) {
	var contact Contact
	err := gs.db.First(&contact, id).Error
	if err != nil {
		return nil, errContactNotFound(id)
	}

	return &contact, err
}

func (gs *GormStore) Update(id int, newName, newEmail string) error {
	contact, err := gs.GetByID(id)
	if err != nil {
		return err
	}
	if newName != "" {
		contact.Name = newName
	}
	if newEmail != "" {
		contact.Email = newEmail
	}
	return gs.db.Save(contact).Error
}

func (gs *GormStore) Delete(id int) error {

	contact, err := gs.GetByID(id)
	if err != nil {
		return err
	}

	if contact == nil {
		return errContactNotFound(id)
	}
	return gs.db.Delete(&Contact{}, id).Error
}
