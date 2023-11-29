package configs

import (
	"sync"

	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils/worker"
	"github.com/hibiken/asynq"
	"github.com/spf13/viper"
)

var taskDistributorInstance worker.TaskDistributor
var taskDistributorOnce sync.Once

func GetTaskDistributor() worker.TaskDistributor {
	taskDistributorOnce.Do(func() {
		redisOpt := asynq.RedisClientOpt{Addr: "payment_service_queue:6379", Password: viper.Get("REDIS_PASSWORD").(string)}
		taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

		taskDistributorInstance = taskDistributor
	})

	return taskDistributorInstance
}
