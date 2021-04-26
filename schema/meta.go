package schema

// Meta represents metadata for a response from the Hetzner DNS API
type Meta struct {
	Pagination struct {
		Page         int `json:"page"`
		PerPage      int `json:"per_page"`
		PreviousPage int `json:"previous_page"`
		NextPage     int `json:"next_page"`
		LastPage     int `json:"last_page"`
		TotalEntries int `json:"total_entries"`
	} `json:"pagination"`
}
