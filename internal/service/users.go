package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/xeoncross/project-templates/database/db"
)

// The users service allows us to separate our business logic from whatever
// implementations we choose. For example, the http server endpoints use this.

type Users interface {
	GetUserByEmail(ctx context.Context, email string) (*db.User, error)
	GetUsers(ctx context.Context) ([]*db.User, error)
	InsertUser(ctx context.Context, u db.User) (int64, error)
}

// concret class user implements the userservice
var _ Users = (*User)(nil)

type User struct {
	DB db.Querier
}

func (u *User) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	users, err := u.DB.GetUserByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, nil // do not require SQL knowledge by the caller
	}
	return users, nil
}

func (u *User) GetUsers(ctx context.Context) ([]*db.User, error) {
	return u.DB.GetUsers(ctx)
}

func (u *User) InsertUser(ctx context.Context, user db.User) (int64, error) {
	params := db.InsertUserParams{Name: user.Name, Email: user.Email}
	return u.DB.InsertUser(ctx, params)
}
