package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GmemberModel = (*customGmemberModel)(nil)

type (
	// GmemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGmemberModel.
	GmemberModel interface {
		gmemberModel
		FindAllGroupByUserId(ctx context.Context, id int64) ([]*Gmember, error)
		FindAllUserByGroupId(ctx context.Context, id int64) ([]*Gmember, error)
		QuietGroup(ctx context.Context, uid int64, gid int64) error
		withSession(session sqlx.Session) GmemberModel
	}

	customGmemberModel struct {
		*defaultGmemberModel
	}
)

func (m *customGmemberModel) QuietGroup(ctx context.Context, uid int64, gid int64) error {
	query := fmt.Sprintf("delete from %s where `uid` = ? and `gid` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, uid, gid)
	return err
}

func (m *customGmemberModel) FindAllGroupByUserId(ctx context.Context, id int64) ([]*Gmember, error) {
	query := fmt.Sprintf("select %s from %s where `uid` = ?", gmemberRows, m.table)
	var resp []*Gmember
	err := m.conn.QueryRowsCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customGmemberModel) FindAllUserByGroupId(ctx context.Context, id int64) ([]*Gmember, error) {
	query := fmt.Sprintf("select %s from %s where `gid` = ?", gmemberRows, m.table)
	var resp []*Gmember
	err := m.conn.QueryRowsCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewGmemberModel returns a model for the database table.
func NewGmemberModel(conn sqlx.SqlConn) GmemberModel {
	return &customGmemberModel{
		defaultGmemberModel: newGmemberModel(conn),
	}
}

func (m *customGmemberModel) withSession(session sqlx.Session) GmemberModel {
	return NewGmemberModel(sqlx.NewSqlConnFromSession(session))
}
