package model

import (
	"context"

	"myapp/util"
)

type userCtxKeyType string

var userCtxKey = userCtxKeyType("user")

func SetUserCtx(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

func GetUserCtx(ctx context.Context) (*User, error) {
	v, ok := ctx.Value(userCtxKey).(*User)
	if !ok {
		return nil, util.NewBadRequestError("not authenticated")
	}

	return v, nil
}
