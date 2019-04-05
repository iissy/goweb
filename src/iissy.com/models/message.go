package models

// Uploador 上传消息
type Uploador struct {
	Success bool   `json:"ok"`
	Message string `json:"msg"`
	Path    string `json:"data"`
}

// Author 未授权信息
type Author struct {
	Success bool   `json:"ok"`
	Message string `json:"msg"`
}
