package models

import (

	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	Sort       string      `json:"sort"`
	Filter	   interface{} `json:"filter"`
	TotalRows  int         `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}


func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "mass desc"
	}
	return p.Sort
}

func (p *Pagination) GetTotalPages() int {
	return (p.GetTotalRows() + p.GetLimit() - 1) / p.GetLimit()
}

func (p *Pagination) GetTotalRows() int {
	return p.TotalRows
}

func paginate(value interface{}, p *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	if p.Filter != nil {
		db = db.Where(p.Filter)
	}
	db.Model(value).Count(&totalRows)
	p.TotalRows = int(totalRows)
	p.TotalPages = p.GetTotalPages()

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort())
	}
}