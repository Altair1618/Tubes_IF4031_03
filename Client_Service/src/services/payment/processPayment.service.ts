import { eq } from "drizzle-orm";
import { db } from "../../configs/drizzle";
import { ProcessPaymentServicePayload } from "../../dto/payment/processPayment.dto";
import { booking } from "../../models/booking";
import { PaymentServiceResponse, ServiceResponse } from "../../types/common";

const processPaymentService = async ({
	paymentUrl,
	jwtToken,
}: ProcessPaymentServicePayload): Promise<ServiceResponse> => {
	try {
		// find booking by payment url
		const bookingData = await db.query.booking.findFirst({
			where: eq(booking.paymentUrl, paymentUrl),
		});

		if (!bookingData) {
			return {
				code: 404,
				message: "booking not found",
			};
		}

		// change booking status
		await db.transaction(async (tx) => {
			await tx
				.update(booking)
				.set({
					status: "PURCHASING",
				})
				.where(eq(booking.id, bookingData.id));

			const response = await fetch(
				`${process.env.PAYMENT_SERVICE_BASE_URL}${paymentUrl}`,
				{
					method: "PATCH",
					headers: {
						Authorization: `Bearer ${jwtToken}`,
					},
				},
			);

			const responseData: PaymentServiceResponse<null> = await response.json();

			if (!response.ok) {
				tx.rollback();
				return {
					code: responseData.code,
					message: responseData.message,
				};
			}
		});

		return {
			code: 200,
			message: "payment processed",
		};
	} catch (e) {
		return {
			code: 500,
			message: "something went wrong, please try again later",
		};
	}
};

export default processPaymentService;
