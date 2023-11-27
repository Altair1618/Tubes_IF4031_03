import { db } from "../../configs/drizzle";
import { and, eq } from "drizzle-orm";
import { ServiceResponse } from "../../types/common";
import { booking } from "../../models/booking";
import { GetBookingServicePayload } from "../../dto/booking/getBooking.dto";
import { BookingData, TicketPriceAndEvent } from "../../types/booking";
import { TicketServiceResponse } from "../../types/common";

const getBookingService = async ({
	userId,
	bookingId,
	jwt,

}: GetBookingServicePayload): Promise<ServiceResponse> => {

	const bookingData = await db.query.booking.findFirst({
		where: and(
			eq(booking.userId, userId),
			eq(booking.id, bookingId)
		)
	});

	if (!bookingData) {
		return {
			code: 404,
			message: "Booking history not found",
		};
	}

	const tiketIds = [bookingData.ticketId]
	
	const params = new URLSearchParams();

	// Append each value of the array as a separate parameter with the same key
	tiketIds.forEach(tiketId => {
		params.append('booking_ids[]', tiketId);
	});

	const response = await fetch(`${process.env.TICKET_SERVICE_BASE_URL}/tikets/ids?${params.toString()}`, {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${jwt}`
		},
		credentials: 'include'
	});

	const responseData = await response.json() as TicketServiceResponse<TicketPriceAndEvent>;

	const tiketPricesAndEventNames = responseData.data


	const id = bookingData.id;
	const ticketId = bookingData.ticketId;
	const status = bookingData.status;
	const paymentUrl = bookingData.report;
	const createdAt = bookingData.createdAt;
	const eventName = tiketPricesAndEventNames[ticketId].eventName;
	const price = tiketPricesAndEventNames[ticketId].price;

	const data: BookingData = {
		id,
		ticketId,
		status,
		paymentUrl,
		createdAt,
		eventName,
		price
	}	

	return {
		code: 200,
		message: "Success",
		data: data
	}
};

export default getBookingService;
