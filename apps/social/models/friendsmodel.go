package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ FriendsModel = (*customFriendsModel)(nil)

type (
	// FriendsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendsModel.
	FriendsModel interface {
		friendsModel
		Inserts(ctx context.Context, session sqlx.Session, data ...*Friends) (sql.Result, error)
		FindByUidAndFid(ctx context.Context, uid, fid string) (*Friends, error)
		ListByUserid(ctx context.Context, userId string) ([]*Friends, error)
	}

	customFriendsModel struct {
		*defaultFriendsModel
	}
)

func (c customFriendsModel) Inserts(ctx context.Context, session sqlx.Session, data ...*Friends) (sql.Result, error) {
	var (
		sql  strings.Builder
		args []any
	)

	if len(data) == 0 {
		return nil, nil
	}

	// insert into tablename values(数据), (数据)
	sql.WriteString(fmt.Sprintf("insert into %s (%s) values ", c.table, friendsRowsExpectAutoSet))

	for i, v := range data {
		sql.WriteString("(?, ?, ?, ?, ?)")
		args = append(args, v.UserId, v.FriendUid, v.Remark, v.AddSource, v.CreatedAt)
		if i == len(data)-1 {
			break
		}

		sql.WriteString(",")
	}

	return session.ExecCtx(ctx, sql.String(), args...)
}

func (c customFriendsModel) FindByUidAndFid(ctx context.Context, uid, fid string) (*Friends, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `friend_uid` = ?", friendsRows, c.table)

	var resp Friends
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, uid, fid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customFriendsModel) ListByUserid(ctx context.Context, userId string) ([]*Friends, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? ", friendsRows, c.table)

	var resp []*Friends
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// NewFriendsModel returns a models for the database table.
func NewFriendsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FriendsModel {
	return &customFriendsModel{
		defaultFriendsModel: newFriendsModel(conn, c, opts...),
	}
}
