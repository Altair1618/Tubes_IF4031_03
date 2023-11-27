package commonStructs

import (
	"time"
)

type EventServicePayload struct {
	EventName string `json:"eventName" form:"eventName" validate:"required"`
	EventTime time.Time `json:"eventTime" form:"eventTime" validate:"required"`
	Location  string `json:"location" form:"location" validate:"required"`
}
