package commonStructs

import (
	"time"

	"github.com/google/uuid"
)

type CreateEventServicePayload struct {
	EventName string    `json:"eventName" form:"eventName" validate:"required"`
	EventTime time.Time `json:"eventTime" form:"eventTime" validate:"required"`
	Location  string    `json:"location" form:"location" validate:"required"`
}

type GetEventsServicePayload struct {
	Query string `query:"query"`
	Page  int    `query:"page" validate:"required"`
}

type UpdateEventServicePayload struct {
	EventName string    `json:"eventName" form:"eventName"`
	EventTime time.Time `json:"eventTime" form:"eventTime"`
	Location  string    `json:"location" form:"location"`
}

type EventDetailResponse struct {
	Id             uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	EventName      string    `json:"eventName" gorm:"not null;column:event_name"`
	EventTime      time.Time `json:"eventTime" gorm:"not null;column:event_time"`
	Location       string    `json:"location" gorm:"not null;column:event_location"`
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at"`
	AvailableSeats int       `json:"availableSeats" gorm:"not null;column:available_seats"`
}
