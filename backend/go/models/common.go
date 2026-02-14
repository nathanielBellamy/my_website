package models

// FilterOptions contains parameters for filtering and sorting list queries.
type FilterOptions struct {
	Page         int
	Limit        int
	ShowInactive bool
	SortField    string
	SortOrder    string // "asc" or "desc"
}

// PaginatedResponse is a generic response structure for list endpoints.
type PaginatedResponse[T any] struct {
	Data  []T `json:"data"`
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
