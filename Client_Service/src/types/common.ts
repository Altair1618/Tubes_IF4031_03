export interface ServiceResponse {
	headers?: HeadersInit;
	code: number;
	message: string;
	data?: unknown;
}
