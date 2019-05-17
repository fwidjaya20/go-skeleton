package contract

// Paginate is a wrapper struct for pagination
type Paginate struct {
	Data     interface{}      `json:"data"`
	Metadata PaginateMetadata `json:"metadata"`
}

// Metadata is struct for additional data
type Metadata struct {
}

// PaginateMetadata is struct for additional paginate information
type PaginateMetadata struct {
	Metadata `json:"-"`
	Total    int64 `json:"total"`
}
