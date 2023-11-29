package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils"
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
		return fmt.Errorf("failed to enqueu task %w", err)
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

	// TODO: simulate 20 percent failure rate
	isFailed := utils.SimulateFailure(20)
	if isFailed {
		log.Info().Str("type", task.Type()).Msg("payment failed, task process succes ❌")
	} else {
		log.Info().Str("type", task.Type()).Msg("payment success, task process success ✅")
	}

	// TODO: hit webhook on ticket service

	return nil
}
