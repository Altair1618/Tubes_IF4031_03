import { z } from "zod";

export const getBookingSchema = z.object({
	userId: z
		.string({
			required_error: "User ID is required",
			invalid_type_error: "User ID is not in a valid type",
		})
		.trim()
		.min(1, {
			message: "User ID is required",
		}),
	bookingId: z
		.string({
			required_error: "User ID is required",
			invalid_type_error: "User ID is not in a valid type",
		})
		.trim()
		.min(1, {
			message: "User ID is required",
		}),
	jwt: z
		.string({
			required_error: "JWT Token is required",
			invalid_type_error: "JWT Token is not in a valid type"
		})
		.trim()
		.min(1, {
			message: "JWT Token is required",
		}),
});

export type GetBookingServicePayload = z.infer<typeof getBookingSchema>;
