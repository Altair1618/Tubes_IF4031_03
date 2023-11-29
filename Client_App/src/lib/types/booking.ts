export enum BookingStatus {
	SUCCESS = "SUCCESS",
	IN_QUEUE = "IN QUEUE", 
	WAITING_FOR_PAYMENT = "WAITING FOR PAYMENT", 
	PURCHASING = "PURCHASING", 
	FAILED = "FAILED"
}

export interface HistoryResponseData {
	id: string,
	ticketId: string,
	status: string,
    paymentUrl: string | null,
	report: string | null,
	createdAt: string,
	seatId: string,
	price: number,
	eventName: string,
	eventTime: string,
	location: string
	totalPage: number
}

export interface CancelBookingResponseData {
	newStatus: string
}