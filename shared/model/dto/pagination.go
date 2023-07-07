package dto

import "math"

type Pagination struct {
	Count     int `json:"count"`
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalPage int `json:"totalPage"`
}

func (p *Pagination) SetPagination(totalData, requestPage, requestSize int) {
	p.Count = totalData
	p.Page = requestPage
	p.PageSize = requestSize
	p.TotalPage = int(math.Ceil(float64(totalData) / float64(requestSize)))
}
