package repositories

import (
	"strings"

	"lms-web-services-main/models/mvc"

	"github.com/LGYtech/lgo"
	"gorm.io/gorm"
)

// ApplyQueryModel: Pagination, Sorting, Filtering ve Searching işlemlerini uygular
func ApplyQueryModel(db *gorm.DB, query *mvc.QueryModel, searchableColumns []string, defaultSorting *mvc.DataSortingOptionItem) (*gorm.DB, *lgo.OperationResult) {
	// Pagination
	if query.PageNumber > 0 && query.RecordsPerPage > 0 {
		db = db.Offset(query.GetSkip()).Limit(query.RecordsPerPage)
	}

	// Sorting
	if len(query.SortingOptions) > 0 {
		for _, sortOption := range query.SortingOptions {
			db = db.Order(sortOption.ToGormOrderString())
		}
	} else if defaultSorting != nil {
		// Varsayılan sıralamayı uygula
		db = db.Order(defaultSorting.ToGormOrderString())
	}

	// Dynamic Filter (Single Filter String)
	if query.Filter != "" {
		parts := strings.Split(query.Filter, "=")
		if len(parts) == 2 {
			field := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			db = db.Where(field+" = ?", value)
		} else {
			return nil, lgo.NewLogicError("Geçersiz filtre formatı. Doğru format: 'field=value'", nil)
		}
	}

	// Searching
	if query.SearchTerm != "" && len(searchableColumns) > 0 {
		var searchQuery strings.Builder
		var searchParams []interface{}

		for i, col := range searchableColumns {
			if i > 0 {
				searchQuery.WriteString(" OR ")
			}
			searchQuery.WriteString(col + " ILIKE ?")
			searchParams = append(searchParams, "%"+query.SearchTerm+"%")
		}

		db = db.Where(searchQuery.String(), searchParams...)
	}

	return db, lgo.NewSuccess(nil)
}
