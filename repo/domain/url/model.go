package url

type URLModel struct {
	Code        string `json:"code"`
	OriginalURL string `json:"original_url"`
	ExpiresAt   int64  `json:"expires_at"` // unix timestamp
	ShortedAt   int64  `json:"shorted_at"` // unix timestamp
}
