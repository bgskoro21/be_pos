package dto

type ApiResponse struct{
	Code	int `json:"statusCode"`
	Data	interface{} `json:"data"`
	Errors	interface{} `json:"errors"`
}