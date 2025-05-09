package mysqluser

import (
	"context"
	"fmt"
	"shop/core/userapp/entity"
	"shop/pkg/errmsg"
	"shop/pkg/richerror"
	"strings"
)

func (db *DB) UpdateUser(ctx context.Context, u entity.User) error {
	const op = "mysqluser.UpdateUser"

	var updateFields []string
	var args []interface{}

	updateFields = append(updateFields, "is_active=?")
	args = append(args, u.IsActive)

	if u.PhoneNumber != "" {
		updateFields = append(updateFields, "phone_number=?", "phone_number_verified=?")
		args = append(args, u.PhoneNumber, 1) // 1 for verified
	}

	if u.Email != "" {
		updateFields = append(updateFields, "email=?", "email_verified=?")
		args = append(args, u.Email, 1) // 1 for verified
	}

	// Ensure at least one field is updated
	if len(updateFields) == 0 {
		return nil
	}

	updateQuery := fmt.Sprintf("UPDATE users SET %s WHERE id=?", strings.Join(updateFields, ", "))
	args = append(args, u.ID)

	_, err := db.conn.Conn().ExecContext(ctx, updateQuery, args...)
	if err != nil {
		return richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}

	return nil
}
