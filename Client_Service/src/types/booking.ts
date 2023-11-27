export interface TicketPriceAndEvent {
	[key: string]: {
		price: number,
    	eventName: string,
		eventTime: string,
		location: string
	};
}

export interface BookingData {
	id: string,
	ticketId: string,
	status: string,
    paymentUrl: string | null,
	createdAt: string,
	eventName: string,
	price: number,
}

export interface BookingDataWithTotalPage extends BookingData {
	totalPage: number
}

export interface CancelBookingData {
	newStatus: string
}