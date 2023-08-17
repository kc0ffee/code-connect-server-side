package models

type AnalyzeResponse struct {
	Type        string  `json:"your_type"`
	Description string  `json:"type_description"`
	Line        int     `json:"code_line"`
	TokenLen    float32 `json:"token_len"`
	TokenCount  int     `json:"token_count"`
	FuncLen     int     `json:"function_len"`
}
