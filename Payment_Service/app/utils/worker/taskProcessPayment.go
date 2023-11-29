package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/configs/database"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
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

	time.Sleep(3 * time.Second)

	db, _ := database.GetGormClient()

	// begin transaction
	tx := db.Begin()

	isFailed := utils.SimulateFailure(20)
	if isFailed {
		log.Info().Str("type", task.Type()).Msg("payment failed, task process succes ❌")
	} else {
		log.Info().Str("type", task.Type()).Msg("payment success, task process success ✅")
	}

	// agent := fiber.Patch(fmt.Sprintf("%s/ticket", viper.Get("TICKET_SERVICE_BASE_URL")))
	agent := fiber.Patch("http://172.22.193.193:3069/api/v1/ticket")
	statusCode, body, errs := agent.Bytes()

	fmt.Println(statusCode)

	if len(errs) > 0 {
		tx.Rollback()
		return fmt.Errorf("webhook request error: %w", errs[0])
	}

	var webhookResponse interface{}
	if err := json.Unmarshal(body, &webhookResponse); err != nil {
		tx.Rollback()
		return fmt.Errorf("webhook request error: %w", err)
	}

	fmt.Println(body)

	if statusCode != fiber.StatusOK {
		tx.Rollback()
		return fmt.Errorf("webhook request error: %w", errs[0])
	}

	tx.Commit()
	// end transaction
	return nil
}
