package models

type CodeDataResponse struct {
	ID        int    `json:"id"`
	Theme     int    `json:"themeId"`
	Lang      string `json:"language"`
	Code      string `json:"code"`
	Timestamp string `json:"timeStamp"`
}

type CodeData struct {
	Theme int    `json:"themeId"`
	Lang  string `json:"language"`
	Code  string `json:"code"`
}
