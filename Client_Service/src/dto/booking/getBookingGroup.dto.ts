import { z } from "zod";
import { isNumeric } from "../../utils/parser";

export const getBookingGroupSchema = z.object({
	userId: z
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
    groupId: z
        .string({
            required_error: "Group ID is required",
        })
        .refine(v => isNumeric(v), {message: "Invalid Group ID"})
});

export type GetBookingGroupServicePayload = z.infer<typeof getBookingGroupSchema>;
