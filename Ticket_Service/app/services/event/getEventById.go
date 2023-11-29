package eventService

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetEventByIdService(id uuid.UUID) utils.ResponseBody {
	db, _ := configs.GetGormClient()

	rows, err := db.Raw(`
		SELECT e.id, e.event_name, e.event_time, e.event_location, e.created_at, e.updated_at, COUNT(t.id) AS available_seats
		FROM events e
		LEFT JOIN tickets t ON e.id = t.event_id
		AND t.status = ?
		GROUP BY e.id
		HAVING e.id = ?
	`, commonStructs.Open, id).Rows()

	if err != nil {
		fmt.Println(err)

		return utils.ResponseBody{
			Code:    500,
			Message: "Error While Fetching Data From Database",
			Data:    nil,
		}
	}

	defer rows.Close()

	event := new(commonStructs.EventDetailResponse)
	for rows.Next() {
		rows.Scan(&event.Id, &event.EventName, &event.EventTime, &event.Location, &event.CreatedAt, &event.UpdatedAt, &event.AvailableSeats)
	}

	return utils.ResponseBody{
		Code:    200,
		Message: "Event Data Fetched Successfully",
		Data:    fiber.Map{"event": event},
	}
}
