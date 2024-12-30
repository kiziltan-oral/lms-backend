package mvc

import (
	"github.com/LGYtech/lgo"
)

// QueryModel: Pagination, Sorting, Searching ve Filtering için tek bir model
type QueryModel struct {
	PageNumber     int                      `json:"pn"`   // Sayfa numarası
	RecordsPerPage int                      `json:"rpp"`  // Sayfa başına kayıt
	SortingOptions []*DataSortingOptionItem `json:"so"`   // Sıralama bilgileri
	Filter         string                   `json:"flt"`  // Tekil string filtre (örn. "Status=Active")
	SearchTerm     string                   `json:"src"`  // Genel arama terimi
	SearchFields   []string                 `json:"srcf"` // Arama yapılacak alanlar
}

// GetSkip: Atlanacak kayıt sayısını hesaplar
func (q *QueryModel) GetSkip() int {
	return (q.PageNumber - 1) * q.RecordsPerPage
}

// Validate: QueryModel'ın geçerliliğini kontrol eder
func (q *QueryModel) Validate() *lgo.OperationResult {
	if q.PageNumber < 1 {
		return lgo.NewLogicError("Geçersiz sayfa numarası. Sayfa numarası 1 veya daha büyük olmalıdır.", nil)
	}
	if q.RecordsPerPage < 1 {
		return lgo.NewLogicError("Gösterilecek veri yoktur.", nil)
	}
	return lgo.NewSuccess(nil)
}

// DataSortingOptionItem: Sıralama için sütun bilgileri
type DataSortingOptionItem struct {
	ColumnName string `json:"cn"` // Sütun adı
	Sorting    int8   `json:"s"`  // 0: ASC, 1: DESC
}

// ToGormOrderString: GORM sıralama stringi oluşturur
func (d *DataSortingOptionItem) ToGormOrderString() string {
	if d.Sorting == 0 {
		return d.ColumnName + " ASC"
	}
	return d.ColumnName + " DESC"
}
