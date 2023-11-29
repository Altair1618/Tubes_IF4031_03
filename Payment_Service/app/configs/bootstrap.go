package configs

import (
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/configs/database"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/models"
	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils/worker"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Bootstrap(app *fiber.App) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	db, err := database.GetGormClient()
	if err != nil {
		log.Fatal().Err(err)
	}

	db.AutoMigrate(&models.Invoice{})
	db.Migrator().DropColumn(&models.Invoice{}, "payment_url")

	_ = GetTaskDistributor()
	redisOpt := asynq.RedisClientOpt{Addr: "payment_service_queue:6379", Password: viper.Get("REDIS_PASSWORD").(string)}
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt)

	go func() {
		log.Info().Msg("start task processor")
		err := taskProcessor.Start()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to start task processor")
		}
	}()
}
