package model

import (
	"errors"
	"gorm.io/gorm/clause"
	"strings"
)

const OrderDirectionASC = "ASC"
const OrderDirectionDESC = "DESC"

type PagingParams struct {
	Offset    int
	Limit     int
	NeedTotal bool
}

type PagingResult struct {
	Offset int
	Limit  int
	Total  int64
}

type OrderParams struct {
	OrderBy        string
	OrderDirection string
}

func (o *OrderParams) ToSqlOrderBy() (c interface{}, err error) {
	if o.OrderBy == "" || o.OrderDirection == "" {
		c = ""
		return
	}
	dir := strings.ToUpper(o.OrderDirection)
	if dir != OrderDirectionASC && dir != OrderDirectionDESC {
		err = errors.New("bad order by")
		return
	}
	return clause.OrderByColumn{Column: clause.Column{Name: o.OrderBy}, Desc: dir == OrderDirectionDESC}, nil
}
