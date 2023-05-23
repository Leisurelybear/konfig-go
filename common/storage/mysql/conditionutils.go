package mysql

import "gorm.io/gorm"

func AppendOrderBy(query *gorm.DB, orderBy string, sort int64) {
	if sort > 0 {
		query.Order(orderBy + " asc")
	} else {
		query.Order(orderBy + " desc")
	}
}
