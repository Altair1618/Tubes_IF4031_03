package worker

import (
	"context"
	"encoding/json"
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/configs/database"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const TaskProcessPayment = "task:process_payment"

func (distributor *RedisTaskDistributor) DistributeTaskProcessPayment(
	ctx context.Context,
	payload *commonStructs.ProcessPaymentServicePayload,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskProcessPayment, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueu task: %w", err)
	}

	log.Info().Str("type", task.Type()).Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")

	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskProcessPayment(_ context.Context, task *asynq.Task) error {
	var payload commonStructs.ProcessPaymentServicePayload

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	db, _ := database.GetGormClient()

	// begin transaction
	tx := db.Begin()

	var requestStatus string
	status, data := processPaymentInit(tx, task, payload)
	if status {
		requestStatus = "SUCCESS"
		tx.Save(&models.Invoice{Id: data.(fiber.Map)["invoiceId"].(models.Invoice).Id, Status: commonStructs.Success})
	} else {
		requestStatus = "FAILED"
	}

	agent := fiber.Patch(fmt.Sprintf("%s/tickets", viper.Get("TICKET_SERVICE_BASE_URL")))
	agent.Set("Authorization", fmt.Sprintf("Bearer %s", payload.JWTToken))
	agent.JSON(fiber.Map{
		"status":    requestStatus,
		"invoiceId": data.(fiber.Map)["invoice"].(models.Invoice).Id,
		"message":   data.(fiber.Map)["message"],
	})
	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		tx.Rollback()
		return fmt.Errorf("webhook request error: %w", errs[0])
	}

	var webhookResponse commonStructs.HttpResponse[interface{}]
	if err := json.Unmarshal(body, &webhookResponse); err != nil {
		tx.Rollback()
		return fmt.Errorf("webhook parse response error: %w", err)
	}

	if statusCode != fiber.StatusOK {
		tx.Rollback()
		return fmt.Errorf(webhookResponse.Message)
	}

	tx.Commit()
	// end transaction
	return nil
}

func processPaymentInit(tx *gorm.DB, task *asynq.Task, payload commonStructs.ProcessPaymentServicePayload) (bool, interface{}) {
	var invoice models.Invoice
	tx.Where("payment_token = ?", payload.PaymentToken).First(&invoice)

	if invoice.Id == uuid.Nil {
		return false, fiber.Map{"message": "invalid payment token", "invoice": invoice}
	} else if invoice.Status == commonStructs.Ongoing {
		return false, fiber.Map{"message": "action not allowed", "invoiceId": invoice}
	}

	// validate token
	token, err := jwt.Parse(payload.JWTToken, func(t *jwt.Token) (interface{}, error) {
		return viper.Get("INVOICE_TOKEN_SECRET").(string), nil
	})
	if err != nil {
		return false, fiber.Map{"message": "invalid payment token", "invoiceId": invoice}
	}
	if !token.Valid {
		return false, fiber.Map{"message": "invalid payment token", "invoiceId": invoice}
	}

	isFailed := utils.SimulateFailure(20)
	if isFailed {
		log.Info().Str("type", task.Type()).Msg("payment failed, task process succes ❌")
		return false, fiber.Map{"message": "payment failed, task process success ❌", "invoiceId": invoice}
	}

	log.Info().Str("type", task.Type()).Msg("payment success, task process success ✅")
	return false, fiber.Map{"message": "payment success, task process success ❌", "invoiceId": invoice}
}
