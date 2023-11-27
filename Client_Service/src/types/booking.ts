export interface TicketPriceAndEvent {
	[key: string]: {
		price: number,
    	eventName: string,
		eventTime: string,
		location: string
	};
}

export interface BookingGroup {
	groupId: number,
	date: string,
	eventName: string,
	totalPrice: number,
	overallStatus: string,
	paymentUrl: string | null,
	totalPage: number
}

export interface BookingGroup {
	groupId: number,
	date: string,
	eventName: string,
	totalPrice: number,
	overallStatus: string,
	paymentUrl: string | null,
	totalPage: number
}