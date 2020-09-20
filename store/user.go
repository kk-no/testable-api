package store

import (
	"context"
	"database/sql"

	"github.com/kk-no/testable-api/database/mysql"
	"github.com/kk-no/testable-api/models"
)

type User struct{ conn *sql.DB }

func NewUser() *User {
	return &User{conn: mysql.Conn}
}

func (u *User) FindByID(ctx context.Context, id string) (*models.User, error) {
	r, err := u.conn.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	user := new(models.User)
	if r.Next() {
		if err := r.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (u *User) Create(ctx context.Context, user *models.User) error {
	_, err := u.conn.QueryContext(ctx, "INSERT INTO users VALUES (?, ?)", user.ID, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(ctx context.Context, user *models.User) error {
	_, err := u.conn.QueryContext(ctx, "UPDATE users SET name = ? WHERE id = ?", user.Name, user.ID)
	if err != nil {
		return err
	}
	return nil
}
