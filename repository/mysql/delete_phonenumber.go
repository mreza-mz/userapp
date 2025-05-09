package mysqluser

import (
	"context"
	"shop/pkg/errmsg"
	"shop/pkg/richerror"
)

func (db *DB) DeletePhoneNumber(ctx context.Context, id uint) error {
	const op = "mysqluser.DeletePhoneNumber"

	const updateQuery = `update users set phone_number=NULL where id = ?`

	_, qErr := db.conn.Conn().ExecContext(ctx, updateQuery, id)

	if qErr != nil {
		return richerror.New(op).WithErr(qErr).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	return nil
}
