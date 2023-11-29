export interface TicketPriceAndEvent {
	tickets: {
		[key: string]: {
			price: number,
			seatId: string,
			eventName: string,
			eventTime: string,
			location: string,
		}	
	}
}

export interface BookingData {
	id: string,
	ticketId: string,
	status: string,
    paymentUrl: string | null,
	report: string | null,
	createdAt: string,
	price: number,
	seatId: string,
	eventName: string,
	eventTime: string,
	location: string,
}

export interface BookingDataWithTotalPage extends BookingData {
	totalPage: number
}

export interface CancelBookingData {
	newStatus: string
}