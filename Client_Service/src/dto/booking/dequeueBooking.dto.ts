import { z } from "zod";

export const dequeueBookingSchema = z.object({
    ticketId: z.string({
        required_error: "Ticket ID is required",
        invalid_type_error: "Ticket ID is not in a valid type",
    }).trim().min(1, {
        message: "Ticket ID is required",
    }).uuid({
        message: "Ticket ID is not in a valid format"
    }),
});

export type DequeueBookingServicePayload = z.infer<typeof dequeueBookingSchema>
