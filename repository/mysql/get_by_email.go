package mysqluser

import (
	"database/sql"
	"errors"
	"shop/core/userapp/entity"
	"shop/pkg/errmsg"
	"shop/pkg/richerror"
)

func (db *DB) GetUserByEmail(email string) (entity.User, error) {
	const op = "mysql.GetUserByEmail"

	const query = `select * from users where email = ?`

	row := db.conn.Conn().QueryRow(query, email)

	user, err := db.scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return entity.User{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return user, nil
}
