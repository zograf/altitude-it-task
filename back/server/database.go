package server

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func createTables() {
	var queries []string
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	queries = append(queries, "DROP TABLE IF EXISTS Users")
	queries = append(queries,
		`CREATE TABLE Users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            lastname VARCHAR(100) NOT NULL,
            birthday DATE NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL,
            is_enabled BOOLEAN DEFAULT FALSE,
            is_deleted BOOLEAN DEFAULT FALSE,
            is_admin BOOLEAN DEFAULT FALSE,
            google_id VARCHAR(100) UNIQUE
        )`)

	for _, query := range queries {
		if _, err := db.Exec(context.Background(), query); err != nil {
			log.Fatalf("failed executing query: %v, error: %v", query, err)
		}
	}
}

func writeUserToDb(user *RegisterDTO, hashedPassword []byte) error {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	query := `INSERT INTO Users (name, lastname, birthday, email, password, is_enabled, is_deleted, is_admin)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var userID int
	err = db.QueryRow(context.Background(), query,
		user.Name,
		user.LastName,
		user.Birthday,
		user.Email,
		string(hashedPassword),
		false,
		false,
		false,
	).Scan(&userID)

	return err
}

func writeAdminToDb() error {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	query := `INSERT INTO Users (email, password, is_admin)
              VALUES ($1, $2, $3) RETURNING id`
	var userID int
	err = db.QueryRow(context.Background(), query,
		ADMIN_EMAIL,
		ADMIN_PASSWORD,
		true,
	).Scan(&userID)

	return err
}

func getUserByEmail(email string) (*User, error) {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	user := User{}

	query := `SELECT id, name, lastname, birthday, email, is_enabled, is_deleted, password
              FROM Users WHERE email = $1`

	err = db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Birthday,
		&user.Email,
		&user.IsEnabled,
		&user.IsDeleted,
		&user.Password,
	)

	return &user, err
}
