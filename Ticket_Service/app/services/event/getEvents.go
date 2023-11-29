package eventService

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetEventsService(payload commonStructs.GetEventsServicePayload) utils.ResponseBody {
	validator := utils.CustomValidator{
		Validator: validator.New(),
	}

	if err := validator.Validate(payload); err != nil {
		return utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: utils.GetValidationErrorMessages(err)[0].Message,
		}
	}

	query := payload.Query
	page := payload.Page

	if page < 1 {
		page = 1
	}

	db, _ := configs.GetGormClient()

	var events []commonStructs.EventDetailResponse

	rows, err := db.Raw(`
		SELECT e.id, e.event_name, e.event_time, e.event_location, e.created_at, e.updated_at, COUNT(t.id) AS available_seats
		FROM events e
		LEFT JOIN tickets t ON e.id = t.event_id
		AND t.status = ?
		WHERE e.event_time > NOW()
		GROUP BY e.id
		HAVING e.event_name LIKE ?
		ORDER BY e.created_at ASC
		LIMIT ? OFFSET ?
	`, commonStructs.Open, fmt.Sprintf("%%%s%%", query), 10, (page-1) * 10).Rows()

	if err != nil {
		fmt.Println(err)

		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: "Error While Fetching Data From Database",
			Data:    nil,
		}
	}

	defer rows.Close()
	for rows.Next() {
		event := new(commonStructs.EventDetailResponse)
		rows.Scan(&event.Id, &event.EventName, &event.EventTime, &event.Location, &event.CreatedAt, &event.UpdatedAt, &event.AvailableSeats)
		events = append(events, *event)
	}

	return utils.ResponseBody{
		Code:    fiber.StatusOK,
		Message: "Events Data Fetched Successfully",
		Data:    fiber.Map{"events": events},
	}
}
