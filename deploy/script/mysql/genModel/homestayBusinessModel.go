package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HomestayBusinessModel = (*customHomestayBusinessModel)(nil)

type (
	// HomestayBusinessModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayBusinessModel.
	HomestayBusinessModel interface {
		homestayBusinessModel
	}

	customHomestayBusinessModel struct {
		*defaultHomestayBusinessModel
	}
)

// NewHomestayBusinessModel returns a model for the database table.
func NewHomestayBusinessModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HomestayBusinessModel {
	return &customHomestayBusinessModel{
		defaultHomestayBusinessModel: newHomestayBusinessModel(conn, c, opts...),
	}
}
