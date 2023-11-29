import Elysia from "elysia";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";
import authMiddleware from "../../middlewares/authMiddleware";
import processPaymentService from "../../services/payment/processPayment.service";
import { ProcessPaymentServicePayload } from "../../dto/payment/processPayment.dto";
import { httpResponse } from "../../utils/httpResponse";

const processPaymentController = new Elysia().use(parseJWTMiddleware).post(
	"/",
	async ({ body, auth: { data } }) => {
		const serviceResponse = await processPaymentService({
			...(body as ProcessPaymentServicePayload),
			jwtToken: data.token,
		});

		return httpResponse(serviceResponse);
	},
	{
		beforeHandle: [
			({ auth: { success, message } }) => {
				return authMiddleware(success, message);
			},
		],
	},
);

export default processPaymentController;
