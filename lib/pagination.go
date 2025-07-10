package pagination

import (
	"math"
	"strconv"
)

type (
	PaginationClientRequest struct {
		Page      string  `json:"page" query:"Page"`
		Size      string  `json:"size" query:"Size"`
		OrderBy   *string `json:"order_by" query:"OrderBy"`
		OrderType *string `json:"order_type" query:"OrderType"`
		Search    *string `json:"search" query:"Search"`
	}
	paginationRequest struct {
		Page   int `json:"page"`
		Size   int `json:"size"`
		Offset int `json:"offset"`
	}
	paginationResponse struct {
		TotalItems  int         `json:"total_items"`
		TotalPages  int         `json:"total_pages"`
		CurrentPage int         `json:"current_page"`
		Items       interface{} `json:"items"`
	}
)

func GetOffset(page string, size string) paginationRequest {
	if len(page) <= 0 {
		page = "1"
	}
	if len(size) > 100 {
		page = "100"
	}
	if len(size) <= 0 {
		size = "10"
	}
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	offsetInt := (pageInt - 1) * sizeInt
	data := paginationRequest{
		Page:   pageInt,
		Size:   sizeInt,
		Offset: offsetInt,
	}
	return data
}

func Data(items interface{}, total_items int, paginationRequest paginationRequest) paginationResponse {
	totalPages := math.Ceil(float64(total_items) / float64(paginationRequest.Size))
	if totalPages <= 0 {
		totalPages = 1
	}
	data := paginationResponse{
		TotalPages:  int(totalPages),
		CurrentPage: paginationRequest.Page,
		TotalItems:  total_items,
		Items:       items,
	}
	return data
}
