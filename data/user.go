package data

import (
	"time"

	up "github.com/upper/db/v4"
)

type User struct {
	ID        int64     `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) Table() string {
	return "users"
}

func (u *User) GetByEmail(email string) (*User, error) {
	collection := upper.Collection(u.Table())

	var newUser User
	res := collection.Find(up.Cond{"email": email})
	err := res.One(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
