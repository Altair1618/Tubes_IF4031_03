export interface ClientServiceResponse<T> {
	code: number;
	message: string;
	data: T;
}

export interface User {
	userId: string;
}

export interface ComboBoxItem<T> {
	value: T,
	label: string
}