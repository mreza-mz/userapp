package mysqluser

import (
	"context"
	"database/sql"
	"errors"
	"shop/core/userapp/entity"
	"shop/pkg/errmsg"
	"shop/pkg/richerror"
)

func (db *DB) GetUserExistByPhoneNumber(ctx context.Context, phoneNumber string) (entity.User, bool, error) {
	const op = "mysql.GetUserExistByPhoneNumber"

	const query = `select * from users where phone_number = ?`

	row := db.conn.Conn().QueryRowContext(ctx, query, phoneNumber)

	user, err := db.scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return user, true, nil
}

func (db *DB) GetUserExistByEmail(ctx context.Context, email string) (entity.User, bool, error) {
	const op = "mysql.GetUserExistByEmail"

	const query = `select * from users where email = ?`

	row := db.conn.Conn().QueryRowContext(ctx, query, email)

	user, err := db.scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return user, true, nil
}
