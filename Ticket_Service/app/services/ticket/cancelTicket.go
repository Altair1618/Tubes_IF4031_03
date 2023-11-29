package ticketService

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/utils"
	"github.com/google/uuid"
)

func CancelTicketService(id uuid.UUID) utils.ResponseBody {
	db, _ := configs.GetGormClient()

	result := db.Where("id = ?", id).Updates(models.Ticket{
			Status: commonStructs.Open,
		},
	)

	if result.Error != nil {
		fmt.Println(result.Error)
		return utils.ResponseBody{
			Code:    500,
			Message: "Error While Updating Data to Database",
			Data:    nil,
		}
	}

	return utils.ResponseBody{
		Code:    200,
		Message: "Tickets Status Set to Open",
	}
}
