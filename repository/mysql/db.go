package mysqluser

import (
	"database/sql"
	"shop/adapter/mysql"
	"shop/core/userapp/entity"
)

type DB struct {
	conn *mysql.DB
}

func New(conn *mysql.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (db *DB) scanUser(scanner mysql.Scanner) (entity.User, error) {
	var user entity.User
	var deletedAt sql.NullTime
	var lastLogin sql.NullTime
	var email sql.NullString
	var phoneNumber sql.NullString
	var password sql.NullString
	var fullname sql.NullString
	var avatar sql.NullString

	err := scanner.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &deletedAt, &lastLogin, &phoneNumber,
		&email, &password, &fullname, &avatar, &user.Role, &user.IsChangedPassword, &user.IsActive)

	if deletedAt.Valid {
		user.DeletedAt = &deletedAt.Time
	}
	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}
	if password.Valid {
		user.Password = password.String
	}
	if phoneNumber.Valid {
		user.PhoneNumber = phoneNumber.String
	}
	if email.Valid {
		user.Email = email.String
	}
	if fullname.Valid {
		user.Fullname = fullname.String
	}
	if avatar.Valid {
		user.Avatar = avatar.String
	}
	return user, err
}
