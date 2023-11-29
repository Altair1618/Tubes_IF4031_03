package paymentService

import (
	"context"
	"fmt"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/common/structs"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/configs"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils/worker"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func ProcessPaymentService(payload commonStructs.ProcessPaymentServicePayload) utils.ResponseBody {
	opts := []asynq.Option{
		asynq.MaxRetry(0),
		asynq.Queue(worker.QueueCritical),
	}

	if taskErr := configs.GetTaskDistributor().DistributeTaskProcessPayment(context.Background(), &payload, opts...); taskErr != nil {
		fmt.Printf("ERROR from service: %s", taskErr.Error())
		log.Err(taskErr).Msg("failed to distribute task")
	}

	return utils.ResponseBody{
		Code:    fiber.StatusOK,
		Message: "payment processed",
	}
}
