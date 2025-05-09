package mysqluser

import (
	"shop/core/userapp/entity"
	"shop/pkg/null"
	"shop/pkg/richerror"
)

func (db *DB) Register(u entity.User) (entity.User, error) {
	const op = "mysql.Register"

	const query = `insert into users(phone_number, email, avatar, role, fullname, password) values(?,?,?,?,?,?)`

	res, err := db.conn.Conn().
		Exec(query, null.NewNullString(u.PhoneNumber), null.NewNullString(u.Email), u.Avatar, u.Role, u.Fullname, u.Password)
	if err != nil {
		return entity.User{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}
