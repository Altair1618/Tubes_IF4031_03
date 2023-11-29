package paymentController

import (
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/common/structs"
	paymentService "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/services/payment"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ProcesPaymentController(c *fiber.Ctx) error {
	payload := new(commonStructs.ProcessPaymentRequest)

	if err := c.ParamsParser(payload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	validator := utils.CustomValidator{
		Validator: validator.New(),
	}

	if err := validator.Validate(payload); err != nil {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: utils.GetValidationErrorMessages(err)[0].Message,
		})
	}

	servicePayload := &commonStructs.ProcessPaymentServicePayload{
		UserId:       c.Locals("userInfo").(commonStructs.JWTPayload).UserId,
		PaymentToken: payload.PaymentToken,
	}

	agent := fiber.Patch("http://ticket_service:3069/api/v1/ticket")
	statusCode, _, errs := agent.Bytes()

	if len(errs) > 0 {
		return utils.CreateResponseBody(c, utils.ResponseBody{
			Code:    400,
			Message: errs[0].Error(),
		})
	}

	fmt.Println(statusCode)

	serviceResponse := paymentService.ProcessPaymentService(*servicePayload)
	return utils.CreateResponseBody(c, serviceResponse)
}
