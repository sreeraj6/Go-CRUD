package models

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ImageUrl     string `json:"image_url"`
	Stock        int    `json:"stock"`
	RecordStatus int    `json:"record_status"`
}
