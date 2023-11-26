import { ServiceResponse } from "../types/common";

export const httpResponse = ({
	headers,
	code,
	message,
	data,
}: ServiceResponse) => {
	return new Response(
		JSON.stringify({
			code,
			message,
			data,
		}),
		{
			//@ts-ignore
			headers,
			status: code,
		},
	);
};
