package store

import (
	"context"
	"database/sql"
	"errors"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"Username"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// represents a store/repository for user data
type UserStore struct {
	// holds a pointer to the sql.DB instance (part of the Go database/sql package)
	// this enables us to perform sql queries against the db
	// by including it in the struct the UsersStore can interact with the db to perform CRUD operations
	db *sql.DB
}

// this func is used to create members in our db
// s *UsersStore indicates that the method is defined on a pointer reciever of the type UsersStore
// meaning that the method can modify the instance it is called on, and allows it to access the db

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (username, password, email)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		// these are the paramaters what we want to insert
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		//the Scan is retreiving the generated feilds (since we dont create these oursleves)
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetById(ctx context.Context, id int64) (*User, error) {
	query := `
	SELECT id, email, username, password, created_at, updated_at
	FROM users
	WHERE id = $1
	`

	user := &User{}

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	// here we are scanning all the feilds created. since it is already done we can just scan for it unlike the create request where they are being inserted
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}
