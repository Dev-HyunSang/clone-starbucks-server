package dto

import "time"

type Status struct {
	Code    int    `json:"code"`
	Stats   string `json:"stats"`
	Message string `json:"message"`
}

type ErrResponse struct {
	Stats       Status    `json:"stats"`
	RespondedAt time.Time `json:"responded_at"`
}
