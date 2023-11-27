package commonStructs

import (
	"time"
)

type DataEventServicePayload struct {
	EventName string `json:"eventName" form:"eventName" validate:"required"`
	EventTime time.Time `json:"eventTime" form:"eventTime" validate:"required"`
	Location  string `json:"location" form:"location" validate:"required"`
}

type GetEventsServicePayload struct {
	Query string `json:"query" form:"query" validate:"required"`
	Page int `json:"page" form:"page" validate:"required"`
}