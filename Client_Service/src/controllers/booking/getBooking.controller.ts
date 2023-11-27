import Elysia from "elysia";
import getBookingService from "../../services/booking/getBooking.service";
import { httpResponse } from "../../utils/httpResponse";
import authMiddleware from "../../middlewares/authMiddleware";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";

const getBookingController = new Elysia().use(parseJWTMiddleware).get(
	"/:id",
	async ({ auth: { data }, params }) => {
		const serviceResponse = await getBookingService({
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
export default getBookingController;
