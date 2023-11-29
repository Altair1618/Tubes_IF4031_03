package paymentController

import (
	"context"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils/worker"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func ProcesPaymentController(c *fiber.Ctx) error {
	taskPayload := &commonStructs.ProcessPaymentServicePayload{
		UserId:       "this is user id",
		PaymentToken: "this is payment token",
	}

	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.Queue(worker.QueueCritical),
	}

	if taskErr := configs.GetTaskDistributor().DistributeTaskProcessPayment(context.Background(), taskPayload, opts...); taskErr != nil {
		log.Err(taskErr).Msg("failed to distribute task")
	}

	return utils.CreateResponseBody(c, utils.ResponseBody{
		Code:    fiber.StatusOK,
		Message: "queue payment success",
	})
}
