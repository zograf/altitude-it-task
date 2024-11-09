package server

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func createTables() {
	var queries []string
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	queries = append(queries, "DROP TABLE IF EXISTS UserConfirmations")
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

	queries = append(queries,
		`CREATE TABLE UserConfirmations (
			id SERIAL PRIMARY KEY,
			uid VARCHAR(255) NOT NULL,
			user_id INTEGER REFERENCES Users(id)
		)`)

	for _, query := range queries {
		if _, err := db.Exec(context.Background(), query); err != nil {
			log.Fatalf("failed executing query: %v, error: %v", query, err)
		}
	}
}

func writeUserToDb(user *RegisterDTO, hashedPassword []byte) (int, error) {
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

	if err != nil {
		fmt.Println(err)
	}

	return userID, err
}

func writeAdminToDb() error {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	query := `INSERT INTO Users (name, lastname, birthday, email, password, is_enabled, is_deleted, is_admin)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var userID int
	err = db.QueryRow(context.Background(), query,
		"",
		"",
		"2000-01-01",
		ADMIN_EMAIL,
		ADMIN_PASSWORD,
		true,
		false,
		true,
	).Scan(&userID)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func getUserByEmail(email string) (*User, error) {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	user := User{}

	query := `SELECT id, email, is_enabled, is_deleted, password, is_admin
	             FROM Users WHERE email = $1`

	err = db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Email,
		&user.IsEnabled,
		&user.IsDeleted,
		&user.Password,
		&user.IsAdmin,
	)
	if err != nil {
		fmt.Println(err)
	}

	return &user, err
}

func writeUidToDb(userId int, uid string) error {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	query := `
        INSERT INTO UserConfirmations (uid, user_id)
        VALUES ($1, $2)
    `
	_, err = db.Exec(context.Background(), query, uid, userId)
	return err
}

func deleteUidFromDb(uid string) error {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	query := `
        DELETE FROM UserConfirmations
        WHERE uid = $1
    `
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec(context.Background(), query, uid)
	return err
}

func getConfirmationByUid(uid string) (*Confirmation, error) {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	conf := Confirmation{}

	query := `SELECT uid, user_id
	             FROM UserConfirmations WHERE uid = $1`

	err = db.QueryRow(context.Background(), query, uid).Scan(
		&conf.Uid,
		&conf.UserId,
	)
	if err != nil {
		fmt.Println(err)
	}

	return &conf, err
}

func enableUserById(userId int) error {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	query := `
        UPDATE Users
        SET is_enabled = true
        WHERE id = $1
    `
	_, err = db.Exec(context.Background(), query, userId)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func writePartialUserToDb(email string) error {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	query := `INSERT INTO Users (name, lastname, birthday, email, password, is_enabled, is_deleted, is_admin)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var userID int
	err = db.QueryRow(context.Background(), query,
		"",
		"",
		"",
		email,
		"",
		true,
		false,
		false,
	).Scan(&userID)

	return err
}

func writeTestUsersToDb() error {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	query := `INSERT INTO Users (name, lastname, birthday, email, password, is_enabled, is_deleted, is_admin)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var userID int
	for i := 0; i < 5; i++ {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("user%d", i+1)), bcrypt.DefaultCost)
		err = db.QueryRow(context.Background(), query,
			fmt.Sprintf("First%d", i+1),
			fmt.Sprintf("Last%d", i+1),
			fmt.Sprintf("2000-0%d-0%d", i+1, i+1),
			fmt.Sprintf("user%d@gmail.com", i+1),
			hashedPassword,
			true,
			false,
			false,
		).Scan(&userID)
		if err != nil {
			fmt.Println(err)
		}
	}

	return err
}

func getUserInfoByEmail(email string) (*User, error) {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	user := User{}

	query := `SELECT id, email, name, lastname, birthday
	             FROM Users WHERE email = $1`

	err = db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.LastName,
		&user.Birthday,
	)
	if err != nil {
		fmt.Println(err)
	}

	return &user, err
}

func updateUserInDb(user UserInfo) error {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	query := `
        UPDATE Users
        SET name = $1, lastname = $2, birthday = $3
        WHERE email = $4
    `

	_, err = db.Exec(context.Background(), query, user.Name, user.LastName, user.Birthday, user.Email)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func updateUserPassword(hashedPassword string, email string) error {
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	query := `
        UPDATE Users
        SET password = $1
        WHERE email = $2
    `

	_, err = db.Exec(context.Background(), query, hashedPassword, email)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
