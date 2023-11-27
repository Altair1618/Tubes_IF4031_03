package commonStructs

import (
	"time"
)

type CreateEventServicePayload struct {
	EventName string    `json:"eventName" form:"eventName" validate:"required"`
	EventTime time.Time `json:"eventTime" form:"eventTime" validate:"required"`
	Location  string    `json:"location" form:"location" validate:"required"`
}

type GetEventsServicePayload struct {
	Query string `json:"query" form:"query"`
	Page  int    `json:"page" form:"page"`
}

type UpdateEventServicePayload struct {
	EventName string    `json:"eventName" form:"eventName"`
	EventTime time.Time `json:"eventTime" form:"eventTime"`
	Location  string    `json:"location" form:"location"`
}
