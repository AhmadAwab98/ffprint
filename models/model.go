package models

// represent expected JSON structure in request body
type BodyRequest struct {
	Path string `json:"path"`
}

// represent JSON structure of response body
type InterimBodyResponse struct {
	Type string `json:"type" redis:"type"`
	Name string `json:"name" redis:"name"`
	Path string `json:"path" redis:"path"`
	Size string `json:"size" redis:"size"`
	LastModified string `json:"lastmodified" redis:"lastmodified"`
	Contents []InterimBodyResponse `json:"contents,omitempty" redis:"contents"`
}

type BodyResponse struct {
	Status string `json:"staus" redis:"staus"`
	Path string `json:"path" redis:"path"`
	Contents []InterimBodyResponse `json:"contents,omitempty" redis:"contents"`
	TFiles int `json:"totalFiles" redis:"totalFiles"`
	TFolders int `json:"totalFolders" redis:"totalFolders"`
	Size int `json:"totalSize" redis:"totalSize"`
}