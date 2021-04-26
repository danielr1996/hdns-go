package schema

// Record represents a DNS record in the Hetzner DNS API
type Record struct {
	Id       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
	ZoneId   string `json:"zone_id,omitempty"`
	Created  string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
}

// RecordResponse represents a Record or multiple Records received from the Hetzner DNS API
type RecordResponse struct {
	Records []Record `json:"records,omitempty"`
	Record  Record   `json:"record,omitempty"`
}

//type RecordsSucess struct {
//	Records []Record `json:"records"`
//}
//
//type RecordSucess struct {
//	Record Record `json:"record"`
//}

// RecordParams represents the query params that can be used to query Records from the Hetzner DNS API
type RecordParams struct {
	ZoneId  string `url:"zone_id,omitempty"`
	Page    string `url:"page,omitempty"`
	PerPage string `url:"per_page,omitempty"`
}
