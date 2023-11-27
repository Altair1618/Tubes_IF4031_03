package commonStructs

import (
	"time"
)

type DataEventServicePayload struct {
	EventName string    `json:"eventName" form:"eventName" validate:"required"`
	EventTime time.Time `json:"eventTime" form:"eventTime" validate:"required"`
	Location  string    `json:"location" form:"location" validate:"required"`
}

type GetEventsServicePayload struct {
	Query string `query:"query"`
	Page  int    `query:"page" validate:"required"`
}
