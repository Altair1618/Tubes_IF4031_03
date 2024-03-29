package configs

import (
	"log"

	"github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/models"
	"github.com/spf13/viper"
)

func Bootstrap() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	db, err := GetGormClient()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Ticket{}, &models.Event{}, &models.TicketInvoiceBooking{})
	// db.Migrator().DropColumn(&models.Event{}, "location")
}
