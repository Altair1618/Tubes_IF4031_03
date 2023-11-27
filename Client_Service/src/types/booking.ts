export interface TicketPriceAndEventName {
	[key: string]: {
		price: number,
    	eventName: string,
	};
}

export interface BookingGroup {
	groupId: number,
	date: string,
	eventName: string,
	totalPrice: number,
	overallStatus: string,
	paymentUrl: string | null
}