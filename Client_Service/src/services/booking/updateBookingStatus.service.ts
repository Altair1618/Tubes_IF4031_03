import { eq } from "drizzle-orm";
import { db } from "../../configs/drizzle";
import { UpdateBookingStatusServicePayload } from "../../dto/booking/updateBookingStatus.dto";
import { booking } from "../../models/booking";
import { ServiceResponse } from "../../types/common";
import dequeueBookingService from "./dequeueBooking.service";

const updateBookingStatusService = async ({
	bookingId,
	status,
	pdf,
}: UpdateBookingStatusServicePayload): Promise<ServiceResponse> => {
	try {
		const bookingData = await db.query.booking.findFirst({
			where: eq(booking.id, bookingId),
		});
		await db.transaction(async (tx) => {
			// TODO: update booking status and save PDF URL
			await tx
				.update(booking)
				.set({
					status,
					report: pdf,
				})
				.where(eq(booking.id, bookingId));

			if (status === "FAILED") {
				// TODO: call release seat endpoint on ticket
				const response = await dequeueBookingService({
					ticketId: bookingData.ticketId,
				});

				tx.rollback();
				return response;
			}
		});

		return {
			code: 200,
			message: "update booking status success",
		};
	} catch (err) {
		return {
			code: 500,
			message: "failed",
		};
	}
};
export default updateBookingStatusService;
