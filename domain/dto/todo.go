package dto

import "time"

type CreateRequest struct {
	Task     string `json:"task"`
	Deadline string `json:"deadline"`
}

type CreateResponse struct {
	ID       uint      `json:"id"`
	Task     string    `json:"task"`
	Deadline time.Time `json:"deadline"`
	Done     bool      `json:"done"`
}
