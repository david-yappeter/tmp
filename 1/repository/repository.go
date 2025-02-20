package repository

import "gorm.io/gorm"

func returnIfErr(err error, ignoredErrors ...error) error {
	for _, er := range ignoredErrors {
		if err == er {
			return nil
		}
	}
	return err
}

func applyPagination(gormDB *gorm.DB, limit *int, page *int) {
	if limit != nil && *limit > 0 {
		gormDB = gormDB.Limit(*limit)
	}

	if page != nil && limit != nil && *page > 0 {
		offset := (*page - 1) * *limit
		gormDB = gormDB.Offset(offset)
	}
}
