package models

// represent expected JSON structure in request body
type BodyRequest struct {
	Path string `json:"path"`
}

// represent JSON structure of response body
type BodyResponse struct {
	Name string `json:"Name" redis:"Name"`
	Contents []BodyResponse `json:"Contents,omitempty" redis:"Contents"`
}