package model

var ClaimsCtxKey = "ClaimsCtxKey"

const (
	DefaultPage    = 1
	DefaultPerPage = 20
)

type CommonParams struct {
	Keyword string `json:"keyword" form:"keyword"`
}

type PaginationParams struct {
	Page    int64 `json:"page" form:"page"`
	PerPage int64 `json:"per_page" form:"per_page"`
}

type PaginationResponse struct {
	Pagination Pagination  `json:"pagination"`
	Data       interface{} `json:"data"`
}

type Pagination struct {
	Page       int64 `json:"page"`
	PerPage    int64 `json:"per_page"`
	TotalItems int64 `json:"total_items"`
}

// GetOffset gets offset to be used on db
func (p *PaginationParams) GetOffset() int64 {
	return (p.Page - 1) * p.PerPage
}

// Validate sets default value
func (p *PaginationParams) Validate() error {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}
	if p.PerPage <= 0 {
		p.PerPage = DefaultPerPage
	}
	return nil
}
