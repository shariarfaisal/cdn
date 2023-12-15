package server

type PaginationDetails struct {
	TotalCount  int         `json:"total_count"`
	TotalPages  int         `json:"total_pages"`
	CurrentPage int         `json:"current_page"`
	Limit       int         `json:"limit"`
	NextPage    int         `json:"next_page"`
	Info        interface{} `json:"info"`
	List        interface{} `json:"list"`
}

func GetPaginationDetails(list interface{}, totalCount int, limit int, page int, info interface{}) *PaginationDetails {
	var totalPages int = 0
	if totalCount > 0 {
		// set default limit, otherwise code will crash if limit is 0
		if limit == 0 {
			limit = 20
		}

		totalPages = totalCount / limit

		if totalCount%limit != 0 {
			totalPages++
		}
	}

	nextPage := page + 1
	if nextPage > totalPages {
		nextPage = 1
	}

	return &PaginationDetails{
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: page,
		Limit:       limit,
		NextPage:    nextPage,
		Info:        info,
		List:        list,
	}
}
