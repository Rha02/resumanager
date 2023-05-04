package models

type Resume struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FileName string `json:"file_name"`
	FileURL  string `json:"file_url"`
	UserID   int    `json:"-"`
	IsMaster bool   `json:"is_master"`
	Size     int    `json:"size"`
}
