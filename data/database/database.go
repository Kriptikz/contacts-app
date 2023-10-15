package database

import (
	"contacts-app/data/models"
)

type IDatabase interface {
	InitDB() error
	CloseDB() error
	GetAllContacts() ([]models.Contact, error)
}

var DB IDatabase

func init() {
	DB = &SQLiteDB{}
}
