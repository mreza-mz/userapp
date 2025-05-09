package mysqluser

import (
	"context"
	"shop/pkg/richerror"
	"time"
)

func (db *DB) UpdateLastLogin(ctx context.Context, userID uint, lastLogin time.Time) error {
	const op = "mysqluser.UpdateUserLastLogin"

	const query = `UPDATE users SET last_login=? WHERE id=?`

	_, err := db.conn.Conn().ExecContext(ctx, query, lastLogin, userID)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
