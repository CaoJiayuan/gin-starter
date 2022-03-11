package database

import (
	"gin-demo/pkg/response"
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type pageRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"perPage"`
}

var ResultFormater = func(data interface{}, meta PageMeta) map[string]interface{} {
	return map[string]interface{}{
		"items": data,
		"total": meta.Total,
	}
}

type Paginator struct {
	page, perPage int
}

type PageMeta struct {
	Total    int64 `json:"total"`
	PerPage  int   `json:"per_page"`
	Page     int   `json:"page"`
	LastPage int   `json:"last_page"`
	From     int   `json:"from"`
	To       int   `json:"to"`
}

func (p *Paginator) Paginate(tx *gorm.DB, targets interface{}) (map[string]interface{}, error) {

	conn := DB()
	var meta PageMeta
	meta.Page = p.page
	meta.PerPage = p.perPage
	meta.From = p.perPage*(p.page-1) + 1

	tx.Statement.Dest = targets
	session := tx.Session(&gorm.Session{})

	e := conn.Table("(?) as `aggregate`", session).Count(&meta.Total).Error

	if e != nil {
		return nil, e
	}
	meta.LastPage = int(math.Ceil(float64(meta.Total) / float64(p.perPage)))

	db := tx.Offset(p.perPage * (p.page - 1)).Limit(p.perPage).Find(targets)

	meta.To = meta.From + int(db.RowsAffected-1)

	return ResultFormater(targets, meta), nil
}

func NewPaginator(page, perPage int) *Paginator {
	if page < 1 {
		page = 1
	}

	return &Paginator{page: page, perPage: perPage}
}

func PageResponse(tx *gorm.DB, targets interface{}, ctx *gin.Context) {
	var req pageRequest
	page := 1
	pageSize := 10
	ctx.Bind(&req)
	if req.Page != 0 {
		page = req.Page
	}
	if req.PageSize != 0 {
		pageSize = req.PageSize
	}

	p := NewPaginator(page, pageSize)
	data, err := p.Paginate(tx, targets)
	if err != nil {
		response.Err(err).End(ctx)
		return
	}

	ctx.JSON(200, data)
}
