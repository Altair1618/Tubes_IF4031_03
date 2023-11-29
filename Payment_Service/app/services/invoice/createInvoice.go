package invoiceService

import (
	"fmt"
	"net/url"
	"time"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func CreateInvoiceService(payload commonStructs.CreateInvoiceServicePayload) utils.ResponseBody {
	validator := utils.CustomValidator{
		Validator: validator.New(),
	}

	if err := validator.Validate(payload); err != nil {
		return utils.ResponseBody{
			Code:    fiber.StatusBadRequest,
			Message: utils.GetValidationErrorMessages(err)[0].Message,
		}
	}

	paymentToken := jwt.NewWithClaims(jwt.SigningMethodHS256, commonStructs.InvoiceTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
		TicketId:         payload.TicketId,
		UserId:           payload.UserId,
	})

	paymentTokenString, err := paymentToken.SignedString([]byte(viper.Get("INVOICE_TOKEN_SECRET").(string)))

	if err != nil {
		fmt.Println(paymentToken)
		fmt.Println(err)

		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: "Error when creating invoice",
			Data:    nil,
		}
	}

	invoice := models.Invoice{
		PaymentToken: paymentTokenString,
		TicketId:     payload.TicketId,
		UserId:       payload.UserId,
	}

	db, _ := configs.GetGormClient()

	result := db.Create(&invoice)

	if result.Error != nil {
		fmt.Println(result.Error)

		return utils.ResponseBody{
			Code:    fiber.StatusInternalServerError,
			Message: "Error when creating invoice",
			Data:    nil,
		}
	} else {
		url := fmt.Sprintf("/payment/%s", url.PathEscape(paymentTokenString))
		return utils.ResponseBody{
			Code:    fiber.StatusOK,
			Message: "Invoice created successfully",
			Data:    fiber.Map{"invoice": invoice, "url": url},
		}
	}
}
