import { db } from "../../configs/drizzle";
import { ServiceResponse, TicketServiceResponse } from "../../types/common";
import { BookingStatus, booking } from "../../models/booking";
import { GetBookingServicePayload } from "../../dto/booking/getBooking.dto";
import { CancelBookingData } from "../../types/booking";
import { and, eq } from "drizzle-orm";
import dequeueBookingService from "./dequeueBooking.service";

const cancelBookingService = async ({
	userId,
	bookingId,
	jwt,
}: GetBookingServicePayload): Promise<ServiceResponse> => {
	const bookingData = await db.query.booking.findFirst({
		where: and(eq(booking.userId, userId), eq(booking.id, bookingId)),
	});

	if (!bookingData) {
		return {
			code: 404,
			message: "Booking history not found",
		};
	}

	if (
		bookingData.status !== BookingStatus.WAITING_FOR_PAYMENT &&
		bookingData.status !== BookingStatus.IN_QUEUE
	) {
		return {
			code: 403,
			message: "Cannot cancel ongoing payment",
		};
	}

	const newStatus = BookingStatus.FAILED;

	try {
		await db.transaction(async (tx) => {
			if (bookingData.status === BookingStatus.WAITING_FOR_PAYMENT) {
				const res = await fetch(
					`${process.env.TICKET_SERVICE_BASE_URL}/api/v1/tickets/${bookingData.ticketId}/status/cancel`,
					{
						method: "PATCH",
						headers: {
							Authorization: `Bearer ${jwt}`,
						},
						credentials: "include",
					},
				);

				const resData: TicketServiceResponse<any> = await res.json();

				if (!res.ok) {
					tx.rollback();
					throw new Error(resData.message);
				}
			}

			await tx
				.update(booking)
				.set({ status: newStatus })
				.where(and(eq(booking.id, bookingId), eq(booking.userId, userId)));

			const response = await dequeueBookingService({
				ticketId: bookingData.ticketId,
			});
			if (response.code !== 200) {
				tx.rollback();
				throw new Error(response.message);
			}
		});

		const cancelData: CancelBookingData = {
			newStatus,
		};

		return {
			code: 200,
			message: "Success",
			data: cancelData,
		};
	} catch (e) {
		console.log(e);

		return {
			code: 500,
			message: "Internal server error",
		};
	}
};

export default cancelBookingService;
