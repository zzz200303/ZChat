package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ ChatlogModel = (*customChatlogModel)(nil)

type (
	// ChatlogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatlogModel.
	ChatlogModel interface {
		chatlogModel
	}

	customChatlogModel struct {
		*defaultChatlogModel
	}
)

// NewChatlogModel returns a model for the mongo.
func NewChatlogModel(url, db, collection string) ChatlogModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customChatlogModel{
		defaultChatlogModel: newDefaultChatlogModel(conn),
	}
}
