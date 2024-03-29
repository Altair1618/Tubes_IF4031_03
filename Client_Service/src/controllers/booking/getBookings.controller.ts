import Elysia from "elysia";
import getBookingsService from "../../services/booking/getBookings.service";
import { httpResponse } from "../../utils/httpResponse";
import authMiddleware from "../../middlewares/authMiddleware";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";

const getBookingsController = new Elysia().use(parseJWTMiddleware).get(
	"/",
	async ({ auth: { data }, query }) => {
		const serviceResponse = await getBookingsService({
			userId: data?.userId,
			page: query.page ?? "1",
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
export default getBookingsController;
