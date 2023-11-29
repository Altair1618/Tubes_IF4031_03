import Elysia from "elysia";
import parseJWTMiddleware from "../../middlewares/parseJWTMiddleware";
import authMiddleware from "../../middlewares/authMiddleware";
import updateBookingStatusService from "../../services/booking/updateBookingStatus.service";
import { UpdateBookingStatusServicePayload } from "../../dto/booking/updateBookingStatus.dto";
import { httpResponse } from "../../utils/httpResponse";

const updateBookingStatusController = new Elysia()
	.use(parseJWTMiddleware)
	.patch(
		"/:id",
		async ({ body, params: { id } }) => {
			const serviceResponse = await updateBookingStatusService({
				...(body as UpdateBookingStatusServicePayload),
				bookingId: id,
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

export default updateBookingStatusController;
