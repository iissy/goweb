package models

// Uploador is hello
type Uploador struct {
	Success bool   `json:"ok"`
	Message string `json:"msg"`
	Path    string `json:"data"`
}
