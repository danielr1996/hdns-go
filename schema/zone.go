package schema

// Zone represents a DNS record in the Hetzner DNS API
type Zone struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Ttl         int      `json:"ttl,omitempty"`
	Registrar   string   `json:"registrar,omitempty"`
	Ns          []string `json:"ns,omitempty"`
	Created     string   `json:"created,omitempty"`
	Verified    string   `json:"verified,omitempty"`
	Modified    string   `json:"modified,omitempty"`
	Project     string   `json:"project,omitempty"`
	Owner       string   `json:"owner,omitempty"`
	Permissions string   `json:"permissions,omitempty"`
	ZoneType    struct {
		Id          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Prices      string `json:"prices,omitempty"`
	} `json:"zone_type,omitempty"`
	Status          string `json:"status,omitempty"`
	Paused          bool   `json:"paused,omitempty"`
	IsSecondaryDns  bool   `json:"is_secondary_dns,omitempty"`
	TxtVerification struct {
		Name  string `json:"name,omitempty"`
		Token string `json:"token,omitempty"`
	} `json:"txt_verification,omitempty"`
	RecordsCount int `json:"records_count,omitempty"`
}

// ZoneResponse represents the Zones received from the Hetzner DNS API with additional Meta data
type ZoneResponse struct {
	Zones []Zone `json:"zones,omitempty"`
	Meta  Meta   `json:"meta,omitempty"`
}

// ZoneParams represents the query params that can be used to query Zones from the Hetzner DNS API
type ZoneParams struct {
	Name       string `url:"name,omitempty"`
	Page       string `url:"page,omitempty"`
	PerPage    string `url:"per_page,omitempty"`
	SearchName string `url:"search_name,omitempty"`
}
