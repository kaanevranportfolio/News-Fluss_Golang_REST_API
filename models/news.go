package models



type News struct {
	Q              string `json:"q,omitempty"`              // Keywords or phrases to search for
	Sources        string `json:"sources,omitempty"`        // Comma-separated source ids
	Domains        string `json:"domains,omitempty"`        // Comma-separated domains
	ExcludeDomains string `json:"excludeDomains,omitempty"` // Comma-separated domains to exclude
	From           string `json:"from,omitempty"`           // Date from (YYYY-MM-DD)
	To             string `json:"to,omitempty"`             // Date to (YYYY-MM-DD)
	Language       string `json:"language,omitempty"`       // Language code (e.g. en)
	SortBy         string `json:"sortBy,omitempty"`         // relevancy, popularity, publishedAt
	PageSize       int    `json:"pageSize,omitempty"`       // 1-100
	Page           int    `json:"page,omitempty"`           // Page number
	SearchIn       string `json:"searchIn,omitempty"`       // title,description,content
}