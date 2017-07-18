package models

// Uploador is hello
type Uploador struct {
	Success bool   `json:"success"`
	Message string `json:"msg"`
	Path    string `json:"file_path"`
}
