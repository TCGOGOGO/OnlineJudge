package models

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	UserId   int64 `db:"user_id"`
	Username string
	Password string
	Email    string
	Phone    string
	School   string
	Motto    string

	TotalLocalSubmit int `db:"total_local_submit"`
	TotalLocalAC     int `db:"total_local_ac"`
	TotalSubmit      int `db:"total_submit"`
	TotalAC          int `db:"total_ac"`

	RegisterTime  time.Time `db:"register_time"`
	LastLoginTime time.Time `db:"last_login_time"`
	IPAddr        string    `db:"ip_addr"`
	Permission    string
	LockStatus    int `db:"lock_status"`
}

type UserModel struct {
	Model
}

func NewUserModel() *UserModel {
	return &UserModel{Model{Table: "Users"}}
}

func (this *UserModel) Insert(tx *sqlx.Tx, user *User) (int64, error) {
	last_insert_id, err := this.InlineInsert(tx, user, nil, []string{"user_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}

func (this *UserModel) QueryById(tx *sqlx.Tx, id int, required []string, excepts []string) (*User, error) {
	user := User{}
	str_fields, err := this.GenerateSelectSQL(user, required, excepts)
	if err != nil {
		return nil, err
	}
	if err := tx.Get(&user, fmt.Sprintf("SELECT %s FROM %s WHERE user_id=?", str_fields, this.Table), id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (this *UserModel) QueryByName(tx *sqlx.Tx, name string, required []string, excepts []string) (*User, error) {

	user := User{}
	str_fields, err := this.GenerateSelectSQL(user, required, excepts)
	if err != nil {
		return nil, err
	}
	if err := tx.Get(&user, fmt.Sprintf("SELECT %s FROM %s WHERE username=? LIMIT 1", str_fields, this.Table), name); err != nil {
		return nil, err
	}
	return &user, nil
}

func (this *UserModel) QueryIdByName(tx *sqlx.Tx, name string) (int64, error) {
	user, err := this.QueryByName(tx, name, []string{"user_id"}, nil)
	if err != nil {
		return 0, err
	}
	if user.UserId == 0 {
		return 0, errors.New("Failed to get user_id")
	}
	return user.UserId, nil
}

func (this *UserModel) Validate(tx *sqlx.Tx, name string, password string) (bool, error) {
	user, err := this.QueryByName(tx, name, []string{"password"}, nil)
	if err != nil {
		return false, err
	}
	if user.Password == password {
		return true, nil
	}
	return false, nil
}
