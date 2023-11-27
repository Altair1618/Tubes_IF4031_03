import Elysia from "elysia";
import cancelBookingService from "../../services/booking/cancelBooking.service";
import { httpResponse } from "../../utils/httpResponse";
import authMiddleware from "../../middlewares/authMiddleware";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";

const cancelBookingController = new Elysia().use(parseJWTMiddleware).patch(
	"/:id/status/cancelled",
	async ({ auth: { data }, params }) => {
		const serviceResponse = await cancelBookingService({
			userId: data?.userId,
			bookingId: params.id,
			jwt: data?.token as string
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
export default cancelBookingController;
