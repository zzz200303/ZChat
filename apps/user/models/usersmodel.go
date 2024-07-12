package models

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		ListByPhone(ctx context.Context, phone string) ([]*Users, error)
		ListByIds(ctx context.Context, ids []string) ([]*Users, error)
		ListByName(ctx context.Context, name string) ([]*Users, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

func (c customUsersModel) ListByPhone(ctx context.Context, phone string) ([]*Users, error) {
	query := fmt.Sprintf("select %s from %s where `phone` like ? ", usersRows, c.table)
	var resp []*Users
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, fmt.Sprint("%", phone, "%"))
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (c customUsersModel) ListByIds(ctx context.Context, ids []string) ([]*Users, error) {
	query := fmt.Sprintf("select %s from %s where `id` like ? ", usersRows, c.table)
	var resp []*Users
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, fmt.Sprint("%", ids, "%"))
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (c customUsersModel) ListByName(ctx context.Context, name string) ([]*Users, error) {
	query := fmt.Sprintf("select %s from %s where `name` like ? ", usersRows, c.table)
	var resp []*Users
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, fmt.Sprint("%", name, "%"))
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}
