import { z } from "zod";

export const requestBookingSchema = z.object({
    ticketId: z.string({
        required_error: "Ticket ID is required",
        invalid_type_error: "Ticket ID is not in a valid type",
    }).trim().min(1, {
        message: "Ticket ID is required",
    }).uuid({
        message: "Ticket ID is not in a valid format"
    }),
    userId: z.string({
        required_error: "User ID is required",
        invalid_type_error: "User ID is not in a valid type",
    }).trim().min(1, {
        message: "User ID is required",
    }),
    jwt: z.string({
        required_error: "JWT is required",
        invalid_type_error: "JWT is not in a valid type",
    }).trim().min(1, {
        message: "JWT is required",
    }),
});

export type RequestBookingServicePayload = z.infer<typeof requestBookingSchema>