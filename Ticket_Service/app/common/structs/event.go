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
	Id             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	EventName      string    `gorm:"not null;column:event_name"`
	EventTime      time.Time `gorm:"not null;column:event_time"`
	Location       string    `gorm:"not null;column:event_location"`
	CreatedAt      time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime;column:updated_at"`
	AvailableSeats int       `gorm:"not null;column:available_seats"`
}
