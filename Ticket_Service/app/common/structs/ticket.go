package ticketStruct

type UpdateStatusServicePayload struct {
	TicketId string `json:"ticketId" form:"ticketId"`
	Status   string `json:"status" form:"status"`
	UserId   string `json:"userId" form:"userId"`
}
