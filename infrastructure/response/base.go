package response

type Meta struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

// APIError
type APIError struct {
	Code    int    `json:"code,omitempty"`
	Type    string `json:"type,omitempty"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type Paginator struct {
	Total  int64 `json:"total"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Link   Link  `json:"links"`
}

type Link struct {
	NextPageUrl string `json:"next_page_url"`
	PrevPageUrl string `json:"prev_page_url"`
}

type MetaPaginator struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Page    Paginator   `json:"page"`
}

type Single struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data, omitempty"`
}

type Paging struct {
	MetaPaginator MetaPaginator `json:"meta"`
	Data          interface{}   `json:"data, omitempty"`
}
