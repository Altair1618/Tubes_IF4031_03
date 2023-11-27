package eventService

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
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

	var events []models.Event

	result := db.Where("event_name LIKE ?", "%"+query+"%").Limit(10).Offset((page - 1) * 10).Find(&events)

	if result.Error != nil {
		fmt.Println(result.Error)

		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: "Error While Fetching Data From Database",
			Data:    nil,
		}
	} else {
		return utils.ResponseBody{
			Code:    fiber.StatusOK,
			Message: "Events Data Fetched Successfully",
			Data:    fiber.Map{"events": events},
		}
	}
}
