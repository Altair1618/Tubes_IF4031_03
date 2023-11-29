import { z } from "zod";

export const updateBookingSchema = z.object({
	status: z.enum(["FAILED", "SUCCESS"]),
	bookingId: z.string(),
	pdf: z
		.string({
			required_error: "pdf is required",
			invalid_type_error: "pdf is not in a valid type",
		})
		.trim()
		.min(1, { message: "pdf is required" }),
});

export type UpdateBookingStatusServicePayload = z.infer<
	typeof updateBookingSchema
>;
