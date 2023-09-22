package helper

import (
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	Offset  int
	PerPage int
	Page    int
}
type Paginator struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
	PerPage   int `json:"per_page"`
}

func PaginationBuilder(perPageString, pageString string) (*Pagination, error) {

	perPage, err := strconv.Atoi(perPageString)
	if err != nil {
		return nil, errors.New("Invalid parameter per_page: not an int")
	}

	page, err := strconv.Atoi(pageString)
	if err != nil {
		return nil, errors.New("Invalid parameter page: not an int")
	}

	showPage := page
	if showPage < 1 {
		showPage = 1
		page = 1
	}

	offset := (page - 1) * perPage
	page = page * perPage

	paginator := Pagination{
		Offset:  offset,
		PerPage: perPage,
		Page:    showPage,
	}
	return &paginator, nil
}
func Paginate(offset, limit int, sort string, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit).Order(sort)
	}
}
