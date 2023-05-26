package model

// 分页参数
type PaginationParams struct {
	Search   string `json:"search"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	Total    int64  `json:"total"`
}

// 查询参数
type QueryParamsWithDate struct {
	Search   string   `json:"search"`
	PageNum  int      `json:"page_num"`
	PageSize int      `json:"page_size"`
	Date     []string `json:"date"`
}
