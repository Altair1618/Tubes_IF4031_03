import { eq, sql } from "drizzle-orm";
import { DequeueBookingServicePayload } from "../../dto/booking/dequeueBooking.dto";
import {
	PaymentServiceResponse,
	ServiceResponse,
	TicketServiceResponse,
} from "../../types/common";
import { db } from "../../configs/drizzle";
import jwt from "jsonwebtoken";
import { booking } from "../../models/booking";

const dequeueBookingService = async ({
	ticketId,
}: DequeueBookingServicePayload): Promise<ServiceResponse> => {
	const query = sql`
        SELECT "id"
        FROM "booking_history"
        WHERE "ticket_id" = ${ticketId}
        AND "status" = 'IN QUEUE'
        ORDER BY "created_at" ASC
    `;

	let result;
	try {
		result = await db.execute(query);
	} catch (e) {
		console.log(e);

		return {
			code: 500,
			message: e.message,
		};
	}

	if (result.rows.length === 0) {
		const url = `${
			process.env.TICKET_SERVICE_BASE_URL
		}/api/v1/tickets/${encodeURIComponent(ticketId)}`;
		const ticketResponse = await fetch(url, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				status: "OPEN",
			}),
		});

		const ticketResponseData: TicketServiceResponse<undefined> =
			await ticketResponse.json();
		if (ticketResponse.status !== 200) {
			return {
				code: 500,
				message: ticketResponseData.message,
			};
		}

		return {
			code: 400,
			message: `There is no booking with ticket ID (${ticketId})`,
		};
	}

	const bookingHistoryId = result.rows[0].id;
	const query = sql`
            UPDATE "booking_history"
            SET "status" = 'WAITING FOR PAYMENT'
            WHERE "id" = ${bookingHistoryId}
        `;

	try {
		await db.execute(query);
	} catch (e) {
		console.log(e);

		return {
			code: 500,
			message: e.message,
		};
	}

	// issue new jwt
	const newAccessToken = jwt.sign(
		{
			userId: result.rows[0].user_id,
			sessionId: "",
		},
		process.env.RSA_PRIVATE_KEY ?? "",
		{ algorithm: "RS256", expiresIn: 60 * 2 },
	);

	const response = await fetch(
		`${process.env.PAYMENT_SERVICE_BASE_URL}/invoice`,
		{
			method: "POST",
			headers: {
				Authorization: `Bearer ${newAccessToken}`,
			},
			body: JSON.stringify({
				ticketId,
			}),
		},
	);

	const responseData: PaymentServiceResponse<any> = await response.json();

	if (!response.ok) {
		return {
			code: responseData.code,
			message: responseData.message,
		};
	}

	// save url
	await db
		.update(booking)
		.set({
			paymentUrl: responseData.data.url,
		})
		.where(eq(booking.id, bookingHistoryId));

	return {
		code: 200,
		message: "Success",
	};
};

export default dequeueBookingService;
