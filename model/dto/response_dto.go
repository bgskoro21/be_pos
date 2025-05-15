package dto

type ApiResponse struct{
	Code	int `json:"statusCode"`
	Data	interface{} `json:"data,omitempty"`
	Errors	interface{} `json:"errors,omitempty"`
}