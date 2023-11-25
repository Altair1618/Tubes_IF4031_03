import { z } from "zod";

export const deleteProfileSchema = z.object({
	userId: z
		.string({
			required_error: "User ID is required",
			invalid_type_error: "User ID is not in a valid type",
		})
		.trim()
		.min(1, {
			message: "User ID is required",
		}),
});

export type DeleteProfileServicePayload = z.infer<typeof deleteProfileSchema>;
