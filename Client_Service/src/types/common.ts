export interface ServiceResponse {
	headers?: HeadersInit;
	code: number;
	message: string;
	data?: unknown;
}

export interface TicketServiceResponse<T> {
	code: number;
	message: string;
	data: T;
}

export interface PaymentServiceResponse<T> {
	code: number;
	message: string;
	data: T;
}
