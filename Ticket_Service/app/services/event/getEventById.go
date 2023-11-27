package eventService

import (
	"fmt"

	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetEventByIdService(id string) utils.ResponseBody {
	db, _ := configs.GetGormClient()

	var event models.Event
	result := db.First(&event, "id = ?", id)

	if result.Error != nil {
		fmt.Println(result.Error)

		return utils.ResponseBody{
			Code:    500,
			Message: "Error While Fetching Data From Database",
			Data:    nil,
		}
	} else {		
		return utils.ResponseBody{
			Code:    200,
			Message: "Event Data Fetched Successfully",
			Data:    fiber.Map{"event": event},
		}
	}
}