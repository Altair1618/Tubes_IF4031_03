package worker

import (
	"context"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/common/structs"
	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskProcessPayment(
		ctx context.Context,
		payload *commonStructs.ProcessPaymentServicePayload,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
	}
}
