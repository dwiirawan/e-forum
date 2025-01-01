package dto

type PaginationListQuestionMeta struct {
	CurrentPage int `json:"currentPage"`
	TotalPage   int `json:"totalPage"`
	TotalData   int `json:"totalData"`
}

type PaginationListQuestionRequest struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Search string `json:"search"`
}
