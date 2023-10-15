package database

import (
	"contacts-app/data/models"
	"database/sql"
)

var _ IDatabase = &SQLiteDB{}

type SQLiteDB struct {
	DB *sql.DB
}

func (s *SQLiteDB) InitDB() error {
	sqlDB, err := sql.Open("sqlite3", "./contacts.db")
	if err != nil {
		return err
	}

	s.DB = sqlDB

	return nil
}

func (s *SQLiteDB) CloseDB() error {
	return s.DB.Close()
}

func (s *SQLiteDB) GetAllContacts() ([]models.Contact, error) {
	rows, err := s.DB.Query("SELECT * FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var c models.Contact
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}

	return contacts, nil
}
