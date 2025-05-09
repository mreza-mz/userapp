package mysqluser

import (
	"context"
	"database/sql"
	"errors"
	"shop/core/userapp/entity"
	"shop/pkg/errmsg"
	"shop/pkg/richerror"
)

func (db *DB) GetUserByID(ctx context.Context, userID uint) (entity.User, error) {
	const op = "mysql.GetUserByID"

	const query = `select * from users where id = ?`

	row := db.conn.Conn().QueryRowContext(ctx, query, userID)
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
